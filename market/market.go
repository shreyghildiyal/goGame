package market

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
	id   ItemId
	name string
}

const DEFAULT_PRICE float32 = 10.0

type Market struct {
	prices     map[ItemId]float32
	id         int
	name       string
	buyOrders  map[ItemId]map[MarketParticipantId]float32
	sellOrders map[ItemId]map[MarketParticipantId]float32
}

func (m *Market) AddBuyOrder(itemId ItemId, buyerId MarketParticipantId, quantity float32) {

	if _, ok := m.prices[itemId]; !ok {
		m.prices[itemId] = DEFAULT_PRICE
	}
	m.buyOrders[itemId][buyerId] = quantity
}

func (m *Market) AddSellOrder(itemId ItemId, sellerId MarketParticipantId, quantity float32) {
	if _, ok := m.prices[itemId]; !ok {
		m.prices[itemId] = DEFAULT_PRICE
	}
	m.sellOrders[itemId][sellerId] = quantity
}

/**
* The idea is that on settlement, the market participants have their inventories and wallets update deltas calculated.
* actioning it is not the market's responsibiity. the Orchestratortakes care of it
* the market decides on the trades and transfers and then updates its prices
 */
func (m *Market) Settle() (map[ItemId]map[MarketParticipantId]float32, map[MarketParticipantId]float32) {

	itemTransfers := make(map[ItemId]map[MarketParticipantId]float32, len(m.prices))
	moneyTransfers := make(map[MarketParticipantId]float32, len(m.prices))
	tradedQuantity := make(map[ItemId]float32, len(m.prices))
	ordersDelta := make(map[ItemId]float32, len(m.prices))

	// update all the transactions that need to happen. money and items
	for itemId, _ := range m.prices {
		totalbuyQuantity := getTotalQuantity(m.buyOrders[itemId])
		totalSellQuantity := getTotalQuantity(m.sellOrders[itemId])

		// var tradedQuantity float32
		if totalSellQuantity < totalbuyQuantity {
			tradedQuantity[itemId] = totalSellQuantity
		} else {
			tradedQuantity[itemId] = totalbuyQuantity
		}

		ordersDelta[itemId] = totalSellQuantity - tradedQuantity[itemId]

		if tradedQuantity[itemId] > 0 {
			// update the itemTransfersMap and moneyTransferMap
			for participantId, bAmount := range m.buyOrders[itemId] {
				itemTransfers[itemId][participantId] += tradedQuantity[itemId] * bAmount / totalbuyQuantity
				moneyTransfers[participantId] -= itemTransfers[itemId][participantId] * m.prices[itemId]

			}

			for participantId, sAmount := range m.sellOrders[itemId] {
				itemTransfers[itemId][participantId] -= tradedQuantity[itemId] * sAmount / totalSellQuantity
				moneyTransfers[participantId] += itemTransfers[itemId][participantId] * m.prices[itemId]

			}

		}
	}

	// after we know what needs to transferred and how much, we need to update the prices in the market

	return itemTransfers, moneyTransfers
}

// There is an assumption here that the quantity is never negative
func getTotalQuantity(orders map[MarketParticipantId]float32) float32 {

	var total float32 = 0
	for _, quant := range orders {
		total += quant
	}
	return total
}
