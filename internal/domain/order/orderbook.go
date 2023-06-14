package order

import "time"

// Money represents any monetary value.
type Money float64

type Order struct {
	Size      Money
	Bid       bool
	Limit     *Limit
	TimeStamp int64
}

// Limit is a group of Orders at a certain price level.
type Limit struct {
	Price       Money
	Orders      []*Order
	TotalVolume Money
}

type OrderBook struct {
	Asks []*Limit // If user wants to sell crypto, they ask.
	Bids []*Limit // If user wants to buy crypto, they bid.
}

func NewLimit(price Money) *Limit {
	return &Limit{
		Price:  price,
		Orders: []*Order{},
	}
}

func NewOrder(bid bool, size Money) *Order {
	return &Order{
		Size:      size,
		Bid:       bid,
		TimeStamp: time.Now().UnixNano(),
	}
}
