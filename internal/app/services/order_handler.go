package services

// TotalVolumeOfBid calculates and returns the total volume of all bids in the order book.
func (ob *CompleteOrderBook) TotalVolumeOfBid() Money {
	totalVolume := Money(0.0)
	for i := 0; i < len(ob.Bids); i++ {
		totalVolume += ob.Bids[i].TotalVolume
	}

	return totalVolume
}

// TotalVolumeOfAsks calculates and returns the total volume of all asks in the order book.
func (ob *CompleteOrderBook) TotalVolumeOfAsks() Money {
	totalVolume := Money(0.0)
	for i := 0; i < len(ob.Asks); i++ {
		totalVolume += ob.Asks[i].TotalVolume
	}

	return totalVolume
}

// IsFilled checks if an order is filled (size equals 0.0) and returns true if it is, false otherwise.
func (o *Order) IsFilled() bool {
	return o.Size == 0.0
}

// Fill fills a given limit order based on the provided order.
// It returns a slice of MatchEngine containing the matches made during the order execution.
func (l *Limit) Fill(o *Order) []MatchEngine {
	var (
		matches        []MatchEngine
		OrdersToDelete []*Order
	)

	for _, order := range l.Orders {
		match := l.FillOrder(order, o)
		matches = append(matches, match)

		l.TotalVolume -= match.SizeFilled

		if order.IsFilled() {
			OrdersToDelete = append(OrdersToDelete, order)
		}

		// end a possible infinity loop.
		if o.IsFilled() {
			break
		}
	}

	for _, order := range OrdersToDelete {
		l.DeleteOrder(order)
	}

	return matches
}

// FillOrder fills an order based on two provided orders.
// It calculates the size filled, updates the orders' sizes, and returns the MatchEngine.
func (l *Limit) FillOrder(a, b *Order) MatchEngine {
	var (
		bid        *Order
		ask        *Order
		SizeFilled Money
	)

	if a.Bid {
		bid = a
		ask = b
	} else {
		bid = b
		ask = a
	}

	// Find the biggest ask or bid to ensure profit is made.
	if a.Size > b.Size {
		a.Size -= b.Size
		SizeFilled = b.Size
		b.Size = 0.0
	} else {
		b.Size -= a.Size
		SizeFilled = a.Size
		a.Size = 0.0
	}

	// Who has the bid or ask, and the size, and at what price the order is executed?
	return MatchEngine{
		Ask:        bid,
		Bid:        ask,
		SizeFilled: SizeFilled,
		Price:      l.Price,
	}
}

// GetLimitByPrice returns the limit at the specified price, or nil if not found.
func (ob *CompleteOrderBook) GetLimitByPrice(price Money) *Limit {
	if limit, found := ob.BidLimits[price]; found {
		return limit
	}
	if limit, found := ob.AskLimits[price]; found {
		return limit
	}
	return nil
}

// FindLimitIndex finds the index of the given order in the limit's orders slice.
func FindLimitIndex(limit *Limit, order *Order) int {
	for index, order := range limit.Orders {
		if order == order {
			return index
		}
	}
	return -1
}

// FindLimit finds the limit at the specified price in a given limit's slice.
// It returns the limit and its index in the slice, or nil and  -1 if not found.
func FindLimit(limits []*Limit, price Money) (*Limit, int) {
	for index, limit := range limits {
		if limit.Price == price {
			return limit, index
		}
	}
	return nil, -1
}

func (ob *CompleteOrderBook) ClearLimit(bid bool, l *Limit) {
	if bid {
		delete(ob.BidLimits, l.Price)
		for i := 0; i < len(ob.Bids); i++ {
			if ob.Bids[i] == l {
				ob.Bids[i] = ob.Bids[len(ob.Bids)-1]
				ob.Bids = ob.Bids[:len(ob.Bids)-1]
			}
		}
	} else {
		delete(ob.AskLimits, l.Price)
		for i := 0; i < len(ob.Asks); i++ {
			if ob.Asks[i] == l {
				ob.Asks[i] = ob.Asks[len(ob.Asks)-1]
				ob.Asks = ob.Asks[:len(ob.Asks)-1]
			}
		}
	}
}

func (bba BuyTheBestAsk) Len() int {
	return len(bba.Limits)
}

func (bba BuyTheBestAsk) Swap(i, j int) {
	bba.Limits[i], bba.Limits[j] = bba.Limits[j], bba.Limits[i]
}

func (bba BuyTheBestAsk) Less(i, j int) bool {
	return bba.Limits[i].Price < bba.Limits[j].Price
}

func (bbb BuyTheBestBid) Len() int {
	return len(bbb.Limits)
}

func (bbb BuyTheBestBid) Swap(i, j int) {
	bbb.Limits[i], bbb.Limits[j] = bbb.Limits[j], bbb.Limits[i]
}

func (bbb BuyTheBestBid) Less(i, j int) bool {
	return bbb.Limits[i].Price > bbb.Limits[j].Price
}

// ----------> For Orders <--------------

func (o Orders) Len() int {
	return len(o)
}

func (o Orders) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func (o Orders) Less(i, j int) bool {
	return o[i].TimeStamp < o[j].TimeStamp
}
