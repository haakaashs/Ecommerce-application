package resources

type CartProduct struct {
	Id                uint64  `json:"id"`
	ProductId         uint64  `json:"product_id"`
	ProductName       string  `json:"product_name"`
	ProductQty        uint    `json:"product_qty"`
	ProductUnitPrice  float64 `json:"product_unit_price"`
	ProductTotalPrice float64 `json:"product_total_price"`
}

type CartResource struct {
	ID           uint64        `json:"id"`
	UserID       uint64        `json:"user_id"`
	Qty          uint          `json:"qty"`
	TotalAmount  float64       `json:"total_amount"`
	CartProducts []CartProduct `json:"cart_products"`
}
