package models

type Cart struct {
	ID           uint64  `gorm:"primaryKey" json:"id,omitempty"`
	UserID       uint64  `gorm:"not null" json:"user_id,omitempty"`
	Qty          uint    `gorm:"not null" json:"qty,omitempty"`
	TotalAmount  float64 `gorm:"not null" json:"total_amount,omitempty"`
	CartProducts []byte  `json:"cart_products"`
}

// gorm:"many2many:cart_products"
func (Cart) TableName() string {
	return "cart"
}

// []resources.CartProduct `gorm:"many2many:cart_products" json:"cart_products,omitempty"`
