package memory

import (
	"context"
	"sync"

	product "sago-sample/internal/product/domain"
)

// ProductRepository is an in-memory implementation of the product.Repository interface
type ProductRepository struct {
	products map[string]*product.Product
	mutex    sync.RWMutex
}

// NewProductRepository creates a new in-memory product repository
func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: make(map[string]*product.Product),
	}
}

// FindByID finds a product by its ID
func (r *ProductRepository) FindByID(ctx context.Context, id product.ProductID) (*product.Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	p, exists := r.products[id.String()]
	if !exists {
		return nil, product.ErrProductNotFound
	}

	return p, nil
}

// FindAll returns all products
func (r *ProductRepository) FindAll(ctx context.Context) ([]*product.Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	products := make([]*product.Product, 0, len(r.products))
	for _, p := range r.products {
		products = append(products, p)
	}

	return products, nil
}

// Save persists a product
func (r *ProductRepository) Save(ctx context.Context, p *product.Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.products[p.ID().String()] = p
	return nil
}

// Delete removes a product
func (r *ProductRepository) Delete(ctx context.Context, id product.ProductID) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.products[id.String()]; !exists {
		return product.ErrProductNotFound
	}

	delete(r.products, id.String())
	return nil
}

// FindByCategory finds products by category ID
func (r *ProductRepository) FindByCategory(ctx context.Context, categoryID product.CategoryID) ([]*product.Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var result []*product.Product
	for _, p := range r.products {
		if p.HasCategory(categoryID) {
			result = append(result, p)
		}
	}

	return result, nil
}
