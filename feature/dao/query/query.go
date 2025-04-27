package query

import (
	"gorm.io/gorm"
)

// Query is the entry point for all queries
type Query struct {
	db      *gorm.DB
	Product Product
}

// Use creates a new Query instance with the given database connection
func Use(db *gorm.DB) *Query {
	q := &Query{
		db: db,
	}

	// Initialize query builders
	q.Product = Product{
		ProductDo: ProductDo{db: db},
		ALL: ProductField{
			ID:          "id",
			Name:        "name",
			Description: "description",
			Price:       "price",
			Currency:    "currency",
			Stock:       "stock",
		},
		ID:          ProductField{ID: "id"},
		Name:        ProductField{Name: "name"},
		Description: ProductField{Description: "description"},
		Price:       ProductField{Price: "price"},
		Currency:    ProductField{Currency: "currency"},
		Stock:       ProductField{Stock: "stock"},
	}

	return q
}
