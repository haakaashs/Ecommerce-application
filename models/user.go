package models

import "time"

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	Email     string    `gorm:"type:varchar(255);not null;unique" json:"email"`
	Role      string    `gorm:"type:enum('admin', 'customer');not null" json:"role"`
	IsActive  bool      `gorm:"type:tinyint(1);not null"`
	Phone     int64     `gorm:"not null;unique" json:"phone"`
	Address   string    `gorm:"type:varchar(255)" json:"address"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Cart      Cart      `gorm:"foreignkey:UserID"`
	Order     []Order   `gorm:"foreignkey:UserID"`
}

func (User) TableName() string {
	return "users"
}
