package models

type Cart struct {
	ID          uint64  `gorm:"primaryKey" json:"id"`
	UserID      uint64  `gorm:"not null" json:"user_id"`
	Products    string  `gorm:"type:json; not null" json:"products"`
	Qty         uint    `gorm:"not null" json:"qty"`
	TotalAmount float64 `gorm:"not null" json:"total_amount"`
}

func (Cart) TableName() string {
	return "cart"
}
