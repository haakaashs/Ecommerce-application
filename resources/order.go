package resources

type OrderStatusUpdate struct {
	OrderId     uint64 `json:"order_id"`
	OrderStatus string `json:"order_status"`
}
