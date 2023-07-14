package order

import (
	"fmt"
	"time"
)

// Money represents any monetary value.
type Money float64

// MatchEngine matches the ask with the bid.
type MatchEngine struct {
	Ask        *Order
	Bid        *Order
	SizeFilled float64
	Price      Money
}

// Order is the container for a buy order content.
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

type CompleteOrderBook struct {
	Asks []*Limit // If user wants to sell crypto, they ask.
	Bids []*Limit // If user wants to buy crypto, they bid.

	AskLimits map[Money]*Limit
	BidLimits map[Money]*Limit
}

func NewOrderBook() *CompleteOrderBook {
	return &CompleteOrderBook{
		Asks:      []*Limit{},
		Bids:      []*Limit{},
		AskLimits: make(map[Money]*Limit),
		BidLimits: make(map[Money]*Limit),
	}
}

func NewLimit(price Money) *Limit {
	return &Limit{
		Price:  price,
		Orders: []*Order{},
	}
}

func (o *Order) String() string {
	return fmt.Sprintf("[size: %.2f]", o.Size)
}

func NewOrder(bid bool, size Money) *Order {
	return &Order{
		Size:      size,
		Bid:       bid,
		TimeStamp: time.Now().UnixNano(),
	}
}

func (l *Limit) AddOrder(o *Order) {
	o.Limit = l
	l.Orders = append(l.Orders, o)
	l.TotalVolume += o.Size
}

func (l *Limit) DeleteOrder(o *Order) {
	for i := 0; i < len(l.Orders); i++ {
		if l.Orders[i] == o {
			l.Orders[i] = l.Orders[len(l.Orders)-1]
			l.Orders = l.Orders[:len(l.Orders)-1]
		}
	}
	o.Limit = nil
	l.TotalVolume -= o.Size

	//TODO: resort the whole orders.
}

func (ob *CompleteOrderBook) PlaceOrder(price Money, o *Order) {

}

func (ob *CompleteOrderBook) add(price Money, o *Order) {

}
