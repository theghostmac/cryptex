package order_test

import (
	"fmt"
	"github.com/theghostmac/cryptex/internal/domain/order"
	"testing"
)

func TestNewLimit(t *testing.T) {
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

	// Optionally, you can print the testLimit value for debugging
	fmt.Println(testLimit)
}

func TestNewOrderBook(t *testing.T) {
	ob := order.NewOrderBook()
	fmt.Println(ob)
}
