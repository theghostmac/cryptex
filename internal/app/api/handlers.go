package api

import (
	"encoding/json"
	"github.com/theghostmac/cryptex/internal/app/services"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CryptoExchangeHandler handles incoming HTTP requests for the cryptoexchange feature.
type CryptoExchangeHandler struct {
	Service *services.CryptoExchangeService
}

func NewCryptoExchangeHandler(service *services.CryptoExchangeService) *CryptoExchangeHandler {
	return &CryptoExchangeHandler{
		Service: service,
	}
}

type TypeOfOrder string

type Order struct {
	Price     services.Money
	Size      services.Money
	Bid       bool
	Timestamp int64
}

type OrderBookData struct {
	Asks []*Order
	Bids []*Order
}

const (
	MarketOrder TypeOfOrder = "MARKET"
	LimitOrder  TypeOfOrder = "LIMIT"
)

// TradeRequest represents the JSON request body for placing a trade.
type TradeRequest struct {
	OrderType TypeOfOrder     `json:"orderType"` // limit or market
	Bid       bool            `json:"bool"`
	Price     services.Money  `json:"price"`
	Size      services.Money  `json:"size"`
	Market    services.Market `json:"market"`
}

// TradeResponse represents the JSON response for a trade.
type TradeResponse struct {
	Message string `json:"message"`
	// Add more fields to the response as needed, e.g., order ID, status, etc.
}

// Trade handles the trade request and responds with the result.
func (exh *CryptoExchangeHandler) Trade(writer http.ResponseWriter, request *http.Request) {
	// Print the raw dataForTrade body and headers for debugging purposes.
	body, _ := io.ReadAll(request.Body)
	log.Printf("Received dataForTrade body: %s", body)
	log.Printf("Received dataForTrade headers: %v", request.Header)

	// Parse the incoming JSON dataForTrade. ✅
	var dataForTrade TradeRequest
	if err := json.NewDecoder(request.Body).Decode(&dataForTrade); err != nil {
		http.Error(writer, "Invalid dataForTrade", http.StatusBadRequest)
		return
	}

	market := services.Market(dataForTrade.Market)
	orderBook := exh.Service.OrderBooks[market]

	placedOrder := services.NewOrder(dataForTrade.Bid, dataForTrade.Size)

	orderBook.PlaceLimitOrder(dataForTrade.Price, placedOrder)

	// write the JSON response.
	response := map[string]interface{}{"msg": "order placed"}
	RespondWithJSON(writer, http.StatusOK, response)

	//// Validate the dataForTrade data (e.g., check if required fields are present).
	//if dataForTrade.OrderType != "bid" && dataForTrade.OrderType != "ask" {
	//	http.Error(writer, "Invalid order type. Should be 'bid' or 'ask'", http.StatusBadRequest)
	//	return
	//}
}

func (exh *CryptoExchangeHandler) GetBook(writer http.ResponseWriter, request *http.Request) error {
	// Retrieve the market parameter from the request URL
	vars := mux.Vars(request)
	market := services.Market(vars["market"])

	// Find the corresponding order book from the Exchange instance
	bookOfOrders, ok := exh.Service.OrderBooks[market]
	if !ok {
		// If market not found, return an error response
		response := map[string]interface{}{"msg": "market not found"}
		RespondWithError(writer, http.StatusBadRequest, response)
	}

	dataFromOrderBook := OrderBookData{
		[]*Order{},
		[]*Order{},
	}
	// Loop through the asks in the order book
	for _, limit := range bookOfOrders.Asks {
		for _, order := range limit.Orders {
			// Process the order as needed
			o := Order{
				Price:     order.Limit.Price,
				Size:      order.Size,
				Bid:       order.Bid,
				Timestamp: order.TimeStamp,
			}
			dataFromOrderBook.Asks = append(dataFromOrderBook.Asks, &o)
			log.Printf("Order: %.2f", order)
		}
	}
	RespondWithJSON(writer, http.StatusOK, dataFromOrderBook)
	return nil
}

// RespondWithJSON is a utility function to respond with a JSON syntax.
func RespondWithJSON(writer http.ResponseWriter, statusCode int, response interface{}) {
	writer.Header().Set("Content-Type", "services/json")
	writer.WriteHeader(statusCode)
	json.NewEncoder(writer).Encode(response)
}

// RespondWithError is a utility function to respond with an error message
func RespondWithError(writer http.ResponseWriter, statusCode int, data interface{}) {
	RespondWithJSON(writer, statusCode, data)
}
