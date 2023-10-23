package models

type Product struct {
	ID              uint64  `gorm:"primary_key" json:"id"`
	Name            string  `gorm:"type:varchar(255);not null" json:"name"`
	Description     string  `gorm:"type:varchar(255);not null" json:"description"`
	ProductCategory string  `gorm:"type:enum('premium', 'regular', 'budget');not null" json:"product_category"`
	Price           float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	InventoryQty    uint    `gorm:"type:int;not null" json:"inventory_qty"`
}

func (Product) TableName() string {
	return "products"
}
