package product

import (
	"errors"
	"time"
)

// Product represents a product entity in the domain
type Product struct {
	id          ProductID
	name        ProductName
	description ProductDescription
	price       Price
	stock       Stock
	categories  []*Category
	createdAt   time.Time
	updatedAt   time.Time
}

// NewProduct creates a new Product entity
func NewProduct(id ProductID, name ProductName, description ProductDescription, price Price, stock Stock) (*Product, error) {
	if id.IsEmpty() {
		return nil, errors.New("product id cannot be empty")
	}

	now := time.Now()
	return &Product{
		id:          id,
		name:        name,
		description: description,
		price:       price,
		stock:       stock,
		categories:  []*Category{},
		createdAt:   now,
		updatedAt:   now,
	}, nil
}

// ID returns the product's ID
func (p *Product) ID() ProductID {
	return p.id
}

// Name returns the product's name
func (p *Product) Name() ProductName {
	return p.name
}

// Description returns the product's description
func (p *Product) Description() ProductDescription {
	return p.description
}

// Price returns the product's price
func (p *Product) Price() Price {
	return p.price
}

// Stock returns the product's stock
func (p *Product) Stock() Stock {
	return p.stock
}

// CreatedAt returns the product's creation time
func (p *Product) CreatedAt() time.Time {
	return p.createdAt
}

// UpdatedAt returns the product's last update time
func (p *Product) UpdatedAt() time.Time {
	return p.updatedAt
}

// UpdateName updates the product's name
func (p *Product) UpdateName(name ProductName) {
	p.name = name
	p.updatedAt = time.Now()
}

// UpdateDescription updates the product's description
func (p *Product) UpdateDescription(description ProductDescription) {
	p.description = description
	p.updatedAt = time.Now()
}

// UpdatePrice updates the product's price
func (p *Product) UpdatePrice(price Price) {
	p.price = price
	p.updatedAt = time.Now()
}

// UpdateStock updates the product's stock
func (p *Product) UpdateStock(stock Stock) {
	p.stock = stock
	p.updatedAt = time.Now()
}

// DecreaseStock decreases the product's stock by the given quantity
func (p *Product) DecreaseStock(quantity uint) error {
	if err := p.stock.Decrease(quantity); err != nil {
		return err
	}
	p.updatedAt = time.Now()
	return nil
}

// IncreaseStock increases the product's stock by the given quantity
func (p *Product) IncreaseStock(quantity uint) {
	p.stock.Increase(quantity)
	p.updatedAt = time.Now()
}

// Categories returns the product's categories
func (p *Product) Categories() []*Category {
	return p.categories
}

// AddCategory adds a category to the product
func (p *Product) AddCategory(category *Category) {
	// Check if category already exists
	for _, c := range p.categories {
		if c.ID() == category.ID() {
			return // Category already exists, do nothing
		}
	}

	p.categories = append(p.categories, category)
	p.updatedAt = time.Now()
}

// RemoveCategory removes a category from the product
func (p *Product) RemoveCategory(categoryID CategoryID) {
	for i, c := range p.categories {
		if c.ID() == categoryID {
			// Remove category by replacing it with the last element and truncating the slice
			p.categories[i] = p.categories[len(p.categories)-1]
			p.categories = p.categories[:len(p.categories)-1]
			p.updatedAt = time.Now()
			return
		}
	}
}

// HasCategory checks if the product belongs to a category
func (p *Product) HasCategory(categoryID CategoryID) bool {
	for _, c := range p.categories {
		if c.ID() == categoryID {
			return true
		}
	}
	return false
}
