package query

import (
	"context"
	"gorm.io/gorm"
	"sago-sample/internal/dao/model"
)

// ProductDo is a query builder for Product
type ProductDo struct {
	db *gorm.DB
}

// Product field names
type ProductField struct {
	ID          string
	Name        string
	Description string
	Price       string
	Currency    string
	Stock       string
}

// Product represents a query builder for Product
type Product struct {
	ProductDo
	ALL         ProductField
	ID          ProductField
	Name        ProductField
	Description ProductField
	Price       ProductField
	Currency    ProductField
	Stock       ProductField
}

// Where adds a where condition to the query

// WithContext sets the context for the query.
func (p *ProductDo) WithContext(ctx context.Context) *ProductDo {
	return &ProductDo{db: p.db.WithContext(ctx)}
}

// Where appends filter conditions to the query builder and returns a new instance.
//func (p *ProductDo) Where(filters ...any) *ProductDo {
//	queryWithFilters := p.db.Where(filters...) // Apply filters to the underlying query
//	return &ProductDo{db: queryWithFilters}    // Return a new instance with the updated query
//}

// First returns the first record that matches the query
func (p *ProductDo) First() (*model.Product, error) {
	var result model.Product
	err := p.db.First(&result).Error
	return &result, err
}

// Find returns all records that match the query
func (p *ProductDo) Find() ([]*model.Product, error) {
	var result []*model.Product
	err := p.db.Find(&result).Error
	return result, err
}

// Save saves a product to the database
func (p *ProductDo) Save(product *model.Product) error {
	return p.db.Save(product).Error
}

// Delete deletes records that match the query
func (p *ProductDo) Delete() (int64, error) {
	result := p.db.Delete(&model.Product{})
	return result.RowsAffected, result.Error
}

// Eq creates an equals condition
//
//	func (field ProductField) Eq(value interface{}) interface{} {
//		return gorm.Expr(string(field)+" = ?", value)
//	}
func Eq(fieldName string, value interface{}) interface{} {
	return gorm.Expr(fieldName+" = ?", value)
}
