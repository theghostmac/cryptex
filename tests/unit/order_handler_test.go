package unit

import (
	"github.com/theghostmac/cryptex/internal/app/services"
	"testing"
)

func TestTotalVolumeOfBid(t *testing.T) {
	orderBook := services.NewOrderBook()
	buyOrder := services.NewOrder(true, 50)
	orderBook.PlaceLimitOrder(18_000, buyOrder)

	// Total volume of bids should be 50.0
	Assert(t, orderBook.TotalVolumeOfBid(), 50.0)
}

func TestTotalVolumeOfAsks(t *testing.T) {
	orderBook := services.NewOrderBook()
	sellOrder := services.NewOrder(false, 20)
	orderBook.PlaceLimitOrder(10_000, sellOrder)

	// Total volume of asks should be 20.0
	Assert(t, orderBook.TotalVolumeOfAsks(), 20.0)
}

func TestIsFilled(t *testing.T) {
	order := services.NewOrder(true, 100)
	// Order is not filled initially
	Assert(t, order.IsFilled(), false)

	order.Size = 0.0
	// Order is now filled
	Assert(t, order.IsFilled(), true)
}

func TestFillOrder(t *testing.T) {
	limit := services.NewLimit(10_000)
	buyOrder := services.NewOrder(true, 100)
	sellOrder := services.NewOrder(false, 50)

	match := limit.FillOrder(buyOrder, sellOrder)
	// buyOrder size reduced by 50 (sellOrder size)
	Assert(t, buyOrder.Size, 50.0)
	// sellOrder size reduced to 0
	Assert(t, sellOrder.Size, 0.0)
	// match should contain the correct Ask, Bid, SizeFilled, and Price
	expectedMatch := services.MatchEngine{
		Ask:        sellOrder,
		Bid:        buyOrder,
		SizeFilled: 50.0,
		Price:      limit.Price,
	}
	Assert(t, match, expectedMatch)
}

// TODO: Add tests for PlaceMarketOrder and other untested functions in services.go
