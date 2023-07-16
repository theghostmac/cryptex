package services

type Money float64

// MatchEngine matches the ask with the bid.
type MatchEngine struct {
	Ask        *Order
	Bid        *Order
	SizeFilled Money // How much is this order being filled for?
	Price      Money
}

// Order is the container for a buy order content.
type Order struct {
	Size      Money
	Bid       bool
	Limit     *Limit
	TimeStamp int64
}

// Limit is a group of Orders at a certain price level with different sizes.
type Limit struct {
	Price       Money
	Orders      Orders
	TotalVolume Money
}

type CompleteOrderBook struct {
	Asks []*Limit // If user wants to sell crypto, they ask.
	Bids []*Limit // If user wants to buy crypto, they bid.

	AskLimits map[Money]*Limit
	BidLimits map[Money]*Limit
}

/*
	The first one is called directional trading and in essence,
	it’s when we buy, wait, and sell.
	If we manage to sell at a price greater than
	the price at which we bought, we make Money.
*/

// Limits houses all limits to sort from.
type Limits []*Limit

// BuyTheBestAsk makes sure the  trader buys only the best ask.
type BuyTheBestAsk struct {
	Limits
}

type BuyTheBestBid struct {
	Limits
}

type Orders []*Order
