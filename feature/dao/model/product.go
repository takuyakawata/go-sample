package model

// Product represents a product in the database
type Product struct {
	ID          string  `gorm:"column:id;primaryKey"`
	Name        string  `gorm:"column:name"`
	Description string  `gorm:"column:description"`
	Price       float64 `gorm:"column:price"`
	Currency    string  `gorm:"column:currency"`
	Stock       int     `gorm:"column:stock"`
	// Add other fields as needed
}

// TableName specifies the table name for the Product model
func (Product) TableName() string {
	return "products"
}
