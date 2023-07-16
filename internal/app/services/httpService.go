package services

type Market string

// CryptoExchangeService ✅ provides methods for interacting with the cryptoexchange.
type CryptoExchangeService struct {
	OrderBooks map[Market]*CompleteOrderBook
}

const (
	MarketETH Market = "ETH"
)

// NewCryptoExchangeService ✅ creates a new CryptoExchangeService instance.
func NewCryptoExchangeService() *CryptoExchangeService {
	bookOfOrders := make(map[Market]*CompleteOrderBook)
	bookOfOrders[MarketETH] = NewOrderBook()

	return &CryptoExchangeService{
		OrderBooks: bookOfOrders,
	}
}