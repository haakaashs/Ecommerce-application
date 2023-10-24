package resources

type CartProduct struct {
	Id                uint64  `json:"-"`
	ProductId         uint64  `json:"product_id" validate:"required"`
	ProductName       string  `json:"product_name" validate:"required"`
	ProductQty        uint    `json:"product_qty" validate:"required"`
	ProductUnitPrice  float64 `json:"product_unit_price"`
	ProductTotalPrice float64 `json:"product_total_price"`
}

type CartResource struct {
	ID           uint64        `json:"id"`
	UserID       uint64        `json:"user_id" validate:"required"`
	Qty          uint          `json:"qty"`
	TotalAmount  float64       `json:"total_amount"`
	Discount     float64       `json:"discount"`
	CartProducts []CartProduct `json:"cart_products"`
}
