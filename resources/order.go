package resources

import "time"

type OrderResource struct {
	ID            uint64        `json:"order_id"`
	UserID        uint64        `json:"user_id" validate:"required"`
	OrderProducts []CartProduct `json:"order_products"`
	OrderStatus   string        `json:"order_status"`
	Dispatched    *time.Time    `json:"dispatched"`
	OrderValue    float64       `json:"order_value"`
	Discount      float64       `json:"discount,omitempty"`
	IsActive      *bool
}

type OrderStatusUpdate struct {
	OrderId     uint64
	OrderStatus string
}
