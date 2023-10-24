package models

type Product struct {
	ID              uint64  `gorm:"primary_key" json:"product_id"`
	Name            string  `gorm:"type:varchar(255);not null" json:"name" validate:"required"`
	Description     string  `gorm:"type:varchar(255);not null" json:"description" validate:"required"`
	ProductCategory string  `gorm:"type:enum('premium', 'regular', 'budget');not null" json:"product_category" validate:"required"`
	Price           float64 `gorm:"type:decimal(10,2);not null" json:"price" validate:"required"`
	InventoryQty    uint    `gorm:"type:int;not null" json:"inventory_qty" validate:"required"`
}

func (Product) TableName() string {
	return "products"
}
