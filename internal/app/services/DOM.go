package services

/*
The DOM displays the aggregated order book data for both bids
(buy orders) and asks (sell orders) at various price levels.
*/

/*
With the DOM implemented, traders will be able to see a
visual representation of the order book's liquidity and
easily observe support and resistance levels,
enabling them to make more informed trading decisions.
*/

// DOM represents the Depth of Market.
type DOM struct {
	Bids []*DOMLevel // list of bid price levels.
	Asks []*DOMLevel // list of ask price levels.
}

// DOMLevel represents the price level in the DOM with the total volume available at that price.
type DOMLevel struct {
	Price  Money // Price level.
	Volume Money // Total volume available at the price level.
}

// UpdateDOM updates the DOM with the latest order book data.
func (dom *DOM) UpdateDOM(ob *CompleteOrderBook) {
	// Clear existing DOM data
	dom.Bids = nil
	dom.Asks = nil

	// Update bid price levels
	for _, limit := range ob.Bids {
		dom.addBidPriceLevel(limit.Price, limit.TotalVolume)
	}

	// Update ask price levels
	for _, limit := range ob.Asks {
		dom.addAskPriceLevel(limit.Price, limit.TotalVolume)
	}
}

// Add a bid price level to the DOM.
func (dom *DOM) addBidPriceLevel(price, volume Money) {
	dom.Bids = append(dom.Bids, &DOMLevel{
		Price:  price,
		Volume: volume,
	})
}

// Add an ask price level to the DOM.
func (dom *DOM) addAskPriceLevel(price, volume Money) {
	dom.Asks = append(dom.Asks, &DOMLevel{
		Price:  price,
		Volume: volume,
	})
}

// Implement functions to handle order updates and keep the DOM up to date.
// For example, when a new order is placed, update the corresponding price level in the DOM.
// When an order is filled or canceled, update the corresponding price level in the DOM.
