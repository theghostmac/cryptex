package main

import (
	"fmt"
	"github.com/theghostmac/cryptex/internal/app/api"
	"github.com/theghostmac/cryptex/internal/app/services"
	"log"
	"net/http"
	"os"
)

func main() {
	// Set logging output to stdout
	log.SetOutput(os.Stdout)

	// Initialize the cryptoexchange application.
	//orderBook := application.NewOrderBook()
	// ✅
	cryptoExchangeService := services.NewCryptoExchangeService()

	// Create a new API handler for the cryptoexchange feature.
	cryptoExchangeHandler := api.NewCryptoExchangeHandler(cryptoExchangeService)

	// Register the API handler to handle incoming requests at the specified endpoint.
	http.HandleFunc("/cryptoexchange/trade", cryptoExchangeHandler.Trade)

	//✅
	http.HandleFunc("/cryptoexchange/trade/place-order", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			// Respond with a success message if the order was placed successfully
			fmt.Fprintln(writer, "Order placed successfully!")
			return
		}
	})

	http.HandleFunc("/cryptoexchange/trade/get-orderbook", func(writer http.ResponseWriter, request *http.Request) {

	})

	// Start the HTTP server and listen for incoming requests.✅
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
