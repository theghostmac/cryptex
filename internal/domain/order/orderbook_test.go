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
	// Assert.
	testLimit.AddOrder(buyOrder)

	fmt.Println(testLimit)
}

func TestNewOrder(t *testing.T) {

}

func TestDeleteOrder(t *testing.T) {
	// Arrange.
	testLimit := order.NewLimit(10_000)
	buyOrderA := order.NewOrder(true, 5)
	buyOrderB := order.NewOrder(true, 6)
	buyOrderC := order.NewOrder(true, 7)
	// Assert.
	testLimit.AddOrder(buyOrderA)
	testLimit.AddOrder(buyOrderB)
	testLimit.AddOrder(buyOrderC)

	// Delete Order
	testLimit.DeleteOrder(buyOrderB)

	fmt.Println(testLimit)
}
