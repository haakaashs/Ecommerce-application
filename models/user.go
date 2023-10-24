package models

type User struct {
	ID       uint64  `gorm:"primary_key" json:"user_id"`
	Name     string  `gorm:"type:varchar(255);not null" json:"name" validate:"required"`
	Password string  `gorm:"type:varchar(255);not null" json:"password" validate:"required"`
	Email    string  `gorm:"type:varchar(255);not null;unique" json:"email" validate:"required"`
	Phone    uint64  `gorm:"not null;unique" json:"phone" validate:"required"`
	Address  string  `gorm:"type:varchar(255)" json:"address"`
	Cart     Cart    `gorm:"foreignkey:UserID" json:"cart"`
	Order    []Order `gorm:"foreignkey:UserID" json:"order,omitempty"`
}

func (User) TableName() string {
	return "users"
}
