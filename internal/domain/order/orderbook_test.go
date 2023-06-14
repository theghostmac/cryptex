package order_test

import (
	"github.com/theghostmac/cryptex/internal/domain/order"
	"testing"
)

func TestNewLimit(t *testing.T) {
	// Arrange.
	testLimit := order.NewLimit(10_000)
	buyOrder := order.NewOrder(true, 5)
}

func TestNewOrder(t *testing.T) {

}
