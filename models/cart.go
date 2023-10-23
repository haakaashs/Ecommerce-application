package models

type Cart struct {
	ID           uint64  `gorm:"primaryKey" json:"id"`
	UserID       uint64  `gorm:"not null" json:"user_id"`
	Qty          uint    `gorm:"not null" json:"qty"`
	TotalAmount  float64 `gorm:"not null" json:"total_amount"`
	CartProducts []byte  `json:"cart_products"`
}

func (Cart) TableName() string {
	return "cart"
}
