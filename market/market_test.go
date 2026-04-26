package market_test

import (
	"testing"

	config "github.com/shreyghildiyal/goGame/configs"
	"github.com/shreyghildiyal/goGame/market"
)

func TestMarketSettle(t *testing.T) {
	initialPrices := map[market.ItemId]float64{
		1: 10,
	}

	cnf := config.Configuration{}
	m := market.NewMarket(1, "test", initialPrices, cnf)

	m.AddBuyOrder(1, 1, 2)
	m.AddSellOrder(1, 1, 2)

	itemTransferMap, moneyTransferMap := m.Settle()

	if len(itemTransferMap) != 1 {
		t.Error("The length of the item transfer map was not 1. It was ", len(itemTransferMap))
	}

	if len(moneyTransferMap) != 1 {
		t.Error("The length of the money transfer map was not 1. It was ", len(moneyTransferMap))
	}
}
