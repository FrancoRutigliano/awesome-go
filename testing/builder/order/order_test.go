package order

import (
	"test-builder-patter/testing/builder/order/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderCalculation(t *testing.T) {
	order := testutil.NewOrderBuilder().WithQuantity(5).WithPrice(20.0).Build()

	total := order.Quantity * int(order.Price)

	assert.Equal(t, 100, total, "El total calculado no es correcto")
}

func TestOrderStatusChange(t *testing.T) {
	order := testutil.NewOrderBuilder().
		WithStatus("shipped").
		Build()

	assert.Equal(t, "shipped", order.Status)
}
