package models

type Cart struct {
	ID          int64   `gorm:"primaryKey" json:"id"`
	UserID      int64   `gorm:"not null" json:"user_id"`
	Products    string  `gorm:"type:json; not null" json:"products"`
	Qty         int     `gorm:"not null" json:"qty"`
	TotalAmount float64 `gorm:"not null" json:"total_amount"`
}

func (Cart) TableName() string {
	return "cart"
}
