package unit

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/theghostmac/cryptex/internal/app/services"
)

// Assert function replaces the testify framework for me.
func Assert(t *testing.T, firstParam, secondParam any) {
	if !reflect.DeepEqual(firstParam, secondParam) {
		t.Errorf("%+v != %+v", firstParam, secondParam)
	}
}

// TestNewLimit is a general test for everything.
func TestNewLimit(t *testing.T) {
	// Arrange.
	testLimit := services.NewLimit(10_000)
	buyOrderA := services.NewOrder(true, 5)
	buyOrderB := services.NewOrder(true, 6)
	buyOrderC := services.NewOrder(true, 7)
	// Assert.
	testLimit.AddOrder(buyOrderA)
	testLimit.AddOrder(buyOrderB)
	testLimit.AddOrder(buyOrderC)

	// Delete Order
	testLimit.DeleteOrder(buyOrderB)

	// Optionally, you can print the testLimit value for debugging
	fmt.Println(testLimit)
}

func Test_NewLimit(t *testing.T) {
	// create a new limit with a price
	price := services.Money(100.0)
	limit := services.NewLimit(price)
	if limit != nil {
		if limit.Price != price {
			t.Errorf("Expected price %.2f, got %.2f", price, limit.Price)
		}
		if len(limit.Orders) != 0 {
			t.Error("Expected empty orders slice, got non-empty one.")
		}
	} else {
		t.Error("NewLimit() returned nil")
	}
}

func TestLimitString(t *testing.T) {
	// Create a new limit with a price and total volume
	price := services.Money(100.0)
	limit := services.NewLimit(price)
	limit.TotalVolume = 50.0

	// Check if LimitString returns the correct formatted string
	expected := "[Price: 100.00 | Volume: 50.00]"
	result := limit.LimitString()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func Test_NewOrderBook(t *testing.T) {
	// create a new order book
	ob := services.NewOrderBook()
	// check if it is initialized successfully.
	if ob == nil {
		t.Error("NewOrderBook() returned nil")
	}
}

func TestOrderString(t *testing.T) {
	// Create a new order with a size
	size := services.Money(0.0)
	order := services.NewOrder(true, size)

	// Check if OrderString returns the correct formatted string
	expected := "[size: 10.00]"
	result := order.OrderString()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestNewOrder(t *testing.T) {
	// Create a new order with a bid and size
	bid := true
	size := services.Money(200.0)
	order := services.NewOrder(bid, size)

	if order != nil {
		if order.Bid != bid {
			t.Errorf("Expected bid status %t, got %t", bid, order.Bid)
		}
		if order.Size != size {
			t.Errorf("Expected size %.2f, got %.2f", size, order.Size)
		}
		if order.TimeStamp <= 0 {
			t.Error("Invalid timestamp for the order")
		}
	} else {
		t.Errorf("error message ..")
	}
}

func TestPlaceLimitOrder(t *testing.T) {
	orderBook := services.NewOrderBook()
	sellOrder := services.NewOrder(false, 10)
	orderBook.PlaceLimitOrder(10_000, sellOrder)
	// check that the size of Ask order book is 1.
	Assert(t, len(orderBook.Asks), 1)

	// Trying test that would fail.
	sellOrder2 := services.NewOrder(false, 20)
	orderBook.PlaceLimitOrder(25_000, sellOrder2)
	Assert(t, len(orderBook.Asks), 2) // change secondParam to 1 to see failing test case.
	// --> fails, nice <---
}

func TestPlaceMarketOrder(t *testing.T) {
	orderBook := services.NewOrderBook()

	// provide liquidity
	sellOrder := services.NewOrder(false, 20)
	orderBook.PlaceLimitOrder(10_000, sellOrder)

	// Trying a test that would fail.
	buyOrder := services.NewOrder(true, 100) // change size to >20 to see failing test case.
	matches := orderBook.PlaceMarketOrder(buyOrder)

	Assert(t, len(matches), 1)
	Assert(t, len(orderBook.Asks), 1)
	Assert(t, orderBook.TotalVolumeOfAsks(), 10.0)
	Assert(t, matches[0].Bid, buyOrder)
	Assert(t, matches[0].SizeFilled, 10.0)
	Assert(t, matches[0].Price, 10_000.0)
	Assert(t, buyOrder.IsFilled(), true)

	fmt.Printf("%+v", matches)
}

func TestPlaceMarketOrderByAWhale(t *testing.T) {
	orderBook := services.NewOrderBook()

	buyOrderA := services.NewOrder(true, 5)
	buyOrderB := services.NewOrder(true, 8)
	buyOrderC := services.NewOrder(true, 10)
	buyOrderD := services.NewOrder(true, 1)

	// Make 3 markets with different price levels
	orderBook.PlaceLimitOrder(5_000, buyOrderC)
	orderBook.PlaceLimitOrder(5_000, buyOrderD)
	orderBook.PlaceLimitOrder(9_000, buyOrderB)
	orderBook.PlaceLimitOrder(10_000, buyOrderA)

	Assert(t, orderBook.TotalVolumeOfBid(), 24.00)

	sellOrder := services.NewOrder(false, 20)
	matches := orderBook.PlaceMarketOrder(sellOrder)

	Assert(t, orderBook.TotalVolumeOfBid(), 4.0)
	Assert(t, len(matches), 3)
	Assert(t, len(orderBook.Bids), 1)

	fmt.Printf("%+v", matches)
}

func TestCancelOrder(t *testing.T) {
	orderBook := services.NewOrderBook()

	buyOrder := services.NewOrder(true, 4)
	orderBook.PlaceLimitOrder(10_000.0, buyOrder)

	Assert(t, orderBook.TotalVolumeOfBid(), 4.0)

	orderBook.CancelOrder(buyOrder)

	Assert(t, orderBook.TotalVolumeOfBid(), 0.0)
}
