package services

import (
	"fmt"
	"sort"
	"time"
)

// NewOrderBook initializes and returns a new CompleteOrderBook instance.
// A CompleteOrderBook holds both the bids and asks lists along with maps to keep track of limits based on their price.
func NewOrderBook() *CompleteOrderBook {
	return &CompleteOrderBook{
		Asks:      []*Limit{},
		Bids:      []*Limit{},
		AskLimits: make(map[Money]*Limit),
		BidLimits: make(map[Money]*Limit),
	}
}

// NewLimit creates and returns a new Limit instance with the specified price.
// A Limit holds a price and a list of orders at that price level.
func NewLimit(price Money) *Limit {
	// Create and return a new Limit instance with the provided price and an empty orders slice.
	return &Limit{
		Price:  price,
		Orders: []*Order{},
	}
}

// LimitString returns a formatted string representation of a Limit instance.
// It displays the size of the order in the format "[size: %.2f]".
func (l *Limit) LimitString() string {
	return fmt.Sprintf("[Price: %.2f | Volume: %.2f]", l.Price, l.TotalVolume)
}

// OrderString returns a formatted string representation of an Order instance.
// It displays the size of the order in the format "[size: %.2f]".
func (o *Order) OrderString() string {
	return fmt.Sprintf("[size: %.2f]", o.Size)
}

// NewOrder creates and returns a new Order instance with the specified bid (true for bid, false for ask) and size.
// An Order represents an individual order in the order book, with its size, bid status, and timestamp.
func NewOrder(bid bool, size Money) *Order {
	return &Order{
		Size:      size,
		Bid:       bid,
		TimeStamp: time.Now().UnixNano(),
	}
}

// AddOrder adds an order to the Limit instance.
// It updates the limit's total volume and appends the order to the orders slice.
func (l *Limit) AddOrder(o *Order) {
	// Set the limit reference in the order.
	o.Limit = l
	// Add the order to the orders slice.
	l.Orders = append(l.Orders, o)
	// Update the total volume of the limit.
	l.TotalVolume += o.Size
}

// DeleteOrder removes an order from the Limit instance.
// It updates the limit's total volume and removes the order from the orders slice.
// After removing the order, it needs to resort the orders to maintain the correct order based on price and timestamp.
func (l *Limit) DeleteOrder(o *Order) {
	// Find the index of the order in the orders slice.
	for i := 0; i < len(l.Orders); i++ {
		if l.Orders[i] == o {
			// Move the last order to the position of the removed order and truncate the slice.
			l.Orders[i] = l.Orders[len(l.Orders)-1]
			l.Orders = l.Orders[:len(l.Orders)-1]
		}
	}

	// Clear the limit reference in the removed order.
	o.Limit = nil
	// Update the total volume of the limit.
	l.TotalVolume -= o.Size

	sort.Sort(l.Orders)
}

// PlaceLimitOrder places a limit order in the order book based on the provided price and order.
// It creates a new limit if it doesn't exist and adds the order to the corresponding bids or asks list.
func (ob *CompleteOrderBook) PlaceLimitOrder(price Money, o *Order) {
	var limit *Limit
	if o.Bid {
		limit = ob.BidLimits[price]
	} else {
		limit = ob.AskLimits[price]
	}
	if limit == nil {
		// Create a new limit if it doesn't exist
		limit = NewLimit(price)
		if o.Bid {
			ob.Bids = append(ob.Bids, limit)
			ob.BidLimits[price] = limit
		} else {
			ob.Asks = append(ob.Asks, limit)
			ob.AskLimits[price] = limit
		}

	}
	limit.AddOrder(o)
}

// SortAsk sorts the asks list in ascending order based on the price and returns it.
func (ob *CompleteOrderBook) SortAsk() []*Limit {
	sort.Sort(BuyTheBestAsk{ob.Asks})
	return ob.Asks
}

// SortBids sorts the bids list in descending order based on the price and returns it.
func (ob *CompleteOrderBook) SortBids() []*Limit {
	sort.Sort(BuyTheBestBid{ob.Bids})
	return ob.Bids
}

// PlaceMarketOrder places a market order in the order book based on the provided price and order.
// It tries to match the order with existing limit orders and fills them accordingly.
// It returns a slice of MatchEngine containing the matches made during the order execution.
func (ob *CompleteOrderBook) PlaceMarketOrder(o *Order) []MatchEngine {
	matches := []MatchEngine{}
	// Order can be bid or ask (buy or sell)
	if o.Bid {
		if o.Size > ob.TotalVolumeOfAsks() {
			panic(fmt.Errorf("not enough volume for market order. \task size [%.2f], market size [%.2f].", ob.TotalVolumeOfAsks(), o.Size))
		}

		for _, limit := range ob.SortAsk() {
			// if it's a bid, check for asks/offers.
			limitMatches := limit.Fill(o)
			matches = append(matches, limitMatches...)

			if len(limit.Orders) == 0 {
				ob.ClearLimit(true, limit)
			}
		}
	} else {
		if o.Size > ob.TotalVolumeOfBid() {
			panic(fmt.Errorf("not enough volume for market order. \task size [%.2f], market size [%.2f].", ob.TotalVolumeOfBid(), o.Size))
		}

		for _, limit := range ob.SortBids() {
			// if it's a bid, check for asks/offers.
			limitMatches := limit.Fill(o)
			matches = append(matches, limitMatches...)

			if len(limit.Orders) == 0 {
				ob.ClearLimit(true, limit)
			}
		}
	}
	return matches
}

func (ob *CompleteOrderBook) CancelOrder(o *Order) {
	limit := o.Limit
	limit.DeleteOrder(o)
}
