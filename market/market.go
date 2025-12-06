package market

import (
	config "github.com/shreyghildiyal/goGame/configs"
)

/*
	I want multiple markets to exist.
	I want instantaneous trasport of goods, because I cant figure out how the planning for future demand etc can work
	Each planet has its own market.
	Trders on each planet can choose to export/import goods outOf/into the market/planet
	amount of goods that can be moved depends on level of trade port
	prices are stored per good per market
	prices move up or down based on the shrtage/surplus of the good in the previous tick
	industries will buy up all the goods for them to fully utilize their empoyees
	industries will immediately put all their goods on the market
	industries get money based on the goods sold
	unsold inventory is returned to the industry
	industry can stockpile upto 5 days of inputs and outputs (assuming max employment)
	traders decide on market orders based on profitability. they will try and trade everything based on profitability and trade weight.
	traders will try to buys/sell from/to each other planet they have access to
	trader access depends on planet heirarchy
	base traders can trade with sector capital
	sector traders can trade with any planet in sector and capitals of neighbouring sectors
	empire traders can trade with any planet in core sector and any sector capital
	inter empire trade happens from "border" sector capitals.
	trade can happen across unclaimed stars to figure out nighbourhood
*/

type ItemId int
type MarketParticipantId int

type TradeItem struct {
	Id   ItemId
	Name string
}

const DEFAULT_PRICE float64 = 10.0

type Market struct {
	prices                   map[ItemId]float64
	id                       int
	name                     string
	buyOrders                map[ItemId]map[MarketParticipantId]float64
	sellOrders               map[ItemId]map[MarketParticipantId]float64
	priceAdjustmentFactor    float64
	adjustmentMinBuyQuantity float64
}

func NewMarket(marketId int, marketName string) Market {

	return Market{
		id:                       marketId,
		name:                     marketName,
		priceAdjustmentFactor:    config.GetConfig().MarketConfs.PriceAdjustmentFactor,
		adjustmentMinBuyQuantity: config.GetConfig().MarketConfs.AdjustmentMinBuyQuantity,
	}
}

func (m *Market) AddBuyOrder(itemId ItemId, buyerId MarketParticipantId, quantity float64) {

	if _, ok := m.prices[itemId]; !ok {
		m.prices[itemId] = DEFAULT_PRICE
	}
	m.buyOrders[itemId][buyerId] = quantity
}

func (m *Market) AddSellOrder(itemId ItemId, sellerId MarketParticipantId, quantity float64) {
	if _, ok := m.prices[itemId]; !ok {
		m.prices[itemId] = DEFAULT_PRICE
	}
	m.sellOrders[itemId][sellerId] = quantity
}

type settlementUpdate struct {
	newPrice         float64
	itemTransferMap  map[MarketParticipantId]float64
	moneyTransferMap map[MarketParticipantId]float64
	itemId           ItemId
}

/**
* The idea is that on settlement, the market participants have their inventories and wallets update deltas calculated.
* actioning it is not the market's responsibiity. the Orchestratortakes care of it
* the market decides on the trades and transfers and then updates its prices
 */
func (m *Market) Settle() (map[ItemId]map[MarketParticipantId]float64, map[MarketParticipantId]float64) {

	itemTransfers := make(map[ItemId]map[MarketParticipantId]float64, len(m.prices))
	moneyTransfers := make(map[MarketParticipantId]float64, len(m.prices))

	itemUpdates := make(chan settlementUpdate)

	// update all the transactions that need to happen. money and items
	for itemId := range m.prices {
		itemId := itemId

		go processItem(itemUpdates, *m, itemId)

	}

	for update := range itemUpdates {
		itemTransfers[update.itemId] = update.itemTransferMap
		for participantId, amount := range update.moneyTransferMap {
			moneyTransfers[participantId] += amount
		}
		m.prices[update.itemId] = update.newPrice
	}

	return itemTransfers, moneyTransfers
}

func processItem(updateChannel chan settlementUpdate, m Market, itemId ItemId) {
	totalbuyQuantity := getTotalQuantity(m.buyOrders[itemId])
	totalSellQuantity := getTotalQuantity(m.sellOrders[itemId])

	var tradedQuantity float64
	// var tradedQuantity float64
	if totalSellQuantity < totalbuyQuantity {
		tradedQuantity = totalSellQuantity
	} else {
		tradedQuantity = totalbuyQuantity
	}

	ordersDelta := totalSellQuantity - tradedQuantity

	updates := settlementUpdate{
		itemTransferMap:  map[MarketParticipantId]float64{},
		moneyTransferMap: map[MarketParticipantId]float64{},
		itemId:           itemId,
	}

	if tradedQuantity > 0 {
		// update the itemTransfersMap and moneyTransferMap
		for participantId, bAmount := range m.buyOrders[itemId] {
			updates.itemTransferMap[participantId] += tradedQuantity * bAmount / totalbuyQuantity
			updates.moneyTransferMap[participantId] -= updates.itemTransferMap[participantId] * m.prices[itemId]

		}

		for participantId, sAmount := range m.sellOrders[itemId] {
			updates.itemTransferMap[participantId] -= tradedQuantity * sAmount / totalSellQuantity
			updates.moneyTransferMap[participantId] += updates.itemTransferMap[participantId] * m.prices[itemId]

		}

	}

	// after we know what needs to transferred and how much, we need to update the prices in the market
	// the idea is that the changes are dependent on the adjustment factor.
	// adjustment factor of 0.05 means that the prices will increase by 5% if the buy orders are double of selle orders
	// similarly the prices will decrease if sell orders are double of buy orders
	// if sell orders are zero, the prices increase only till a stopGap threshold is hit ??
	// alternatively we can have a min demand count below which prices dont increase if there are zero sell orders
	// the second approach seems to be less brittle, if we are sure the demand is going to keep going down with increases prices

	priceAdjustmentFactor := config.GetConfig().MarketConfs.PriceAdjustmentFactor
	adjustmentMinBuyQuantity := config.GetConfig().MarketConfs.AdjustmentMinBuyQuantity
	if totalSellQuantity == 0 {
		if totalbuyQuantity >= adjustmentMinBuyQuantity {
			updates.newPrice = m.prices[itemId] * (1 + priceAdjustmentFactor)
		}
	} else {

		adjustedPriceFactor := priceAdjustmentFactor * ordersDelta / totalSellQuantity
		updates.newPrice = m.prices[itemId] * (1 + adjustedPriceFactor)

	}

	updateChannel <- updates

}

// There is an assumption here that the quantity is never negative
func getTotalQuantity(orders map[MarketParticipantId]float64) float64 {

	var total float64 = 0
	for _, quant := range orders {
		total += quant
	}
	return total
}
