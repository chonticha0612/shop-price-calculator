package shop

type Order struct {
	orders     []OrderItem
	memberCard bool
}
type OrderItem struct {
	name     string
	quantity int
}

func (o Order) MapOrder() map[string]int {
	orderMap := make(map[string]int)
	for _, item := range o.orders {
		orderMap[item.name] += item.quantity
	}
	return orderMap
}
