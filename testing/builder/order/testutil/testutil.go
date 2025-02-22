package testutil

type Order struct {
	ID       int
	Quantity int
	Price    float64
	Status   string
}

// OrderBuilder ayuda a construir órdenes para los tests.
type OrderBuilder struct {
	order Order
}

// NewOrderBuilder crea un nuevo OrderBuilder con valores por defecto.
func NewOrderBuilder() *OrderBuilder {
	return &OrderBuilder{
		order: Order{
			ID:       1,
			Quantity: 1,
			Price:    10.0,
			Status:   "pending",
		},
	}
}

// Métodos para modificar el objeto según sea necesario
func (b *OrderBuilder) WithID(id int) *OrderBuilder {
	b.order.ID = id
	return b
}

func (b *OrderBuilder) WithQuantity(quantity int) *OrderBuilder {
	b.order.Quantity = quantity
	return b
}

func (b *OrderBuilder) WithPrice(price float64) *OrderBuilder {
	b.order.Price = price
	return b
}

func (b *OrderBuilder) WithStatus(status string) *OrderBuilder {
	b.order.Status = status
	return b
}

// Build devuelve la instancia final de Order.
func (b *OrderBuilder) Build() Order {
	return b.order
}
