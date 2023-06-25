package order_test

import (
	"fmt"
	"github.com/theghostmac/cryptex/internal/domain/order"
	"testing"
)

func TestNewLimit(t *testing.T) {
	// Arrange.
	testLimit := order.NewLimit(10_000)
	buyOrder := order.NewOrder(true, 5)

	testLimit.AddOrder(buyOrder)

	fmt.Println(testLimit)
}

func TestNewOrder(t *testing.T) {

}
