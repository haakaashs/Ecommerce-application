package models

import "time"

type Order struct {
	ID          uint64       `gorm:"primary_key" json:"id"`
	UserID      uint64       `gorm:"not null" json:"user_id"`
	Products    string     `gorm:"type:json;not null" json:"products"`
	OrderStatus string     `gorm:"type:enum('placed', 'dispatched', 'completed', 'cancelled');not null" json:"order_status"`
	Dispatched  time.Time  `gorm:"type:datetime" json:"dispatched"`
	OrderValue  float64    `gorm:"type:decimal(10,2);not null" json:"order_value"`
	IsActive    bool       `gorm:"type:tinyint(1);not null"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	ModifiedAt  *time.Time 
}

func (Order) TableName() string {
	return "orders"
}
