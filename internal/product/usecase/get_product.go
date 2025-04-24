package product

import (
	"context"
	"errors"

	domain "sago-sample/internal/product/domain"
)

// GetProductByIDInput represents the input data for getting a product by ID
type GetProductByIDInput struct {
	ID string
}

// GetProductByIDOutput represents the output data after getting a product
type GetProductByIDOutput struct {
	ID          string
	Name        string
	Description string
	Price       uint
	Currency    string
	Stock       uint
	Categories  []CategoryOutput
}

// GetProductByIDUseCase defines the use case for getting a product by ID
type GetProductByIDUseCase struct {
	productService *domain.Service
}

// NewGetProductByIDUseCase creates a new instance of GetProductByIDUseCase
func NewGetProductByIDUseCase(productService *domain.Service) *GetProductByIDUseCase {
	return &GetProductByIDUseCase{
		productService: productService,
	}
}

// Execute runs the use case
func (uc *GetProductByIDUseCase) Execute(ctx context.Context, input GetProductByIDInput) (*GetProductByIDOutput, error) {
	// Create value object
	productID, err := domain.NewProductID(input.ID)
	if err != nil {
		return nil, err
	}

	// Call domain service to get product
	foundProduct, err := uc.productService.GetProductByID(ctx, productID)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	// Map domain entity to output
	// Map categories
	categories := make([]CategoryOutput, 0, len(foundProduct.Categories()))
	for _, c := range foundProduct.Categories() {
		categories = append(categories, CategoryOutput{
			ID:   c.ID().String(),
			Name: c.Name().String(),
		})
	}

	return &GetProductByIDOutput{
		ID:          foundProduct.ID().String(),
		Name:        foundProduct.Name().String(),
		Description: foundProduct.Description().String(),
		Price:       foundProduct.Price().Amount(),
		Currency:    foundProduct.Price().Currency(),
		Stock:       foundProduct.Stock().Quantity(),
		Categories:  categories,
	}, nil
}

// GetAllProductsOutput represents a product in the list of all products
type GetAllProductsOutput struct {
	Products []ProductOutput
}

// ProductOutput represents a product in the output
type ProductOutput struct {
	ID          string
	Name        string
	Description string
	Price       uint
	Currency    string
	Stock       uint
	Categories  []CategoryOutput
}

// GetAllProductsUseCase defines the use case for getting all products
type GetAllProductsUseCase struct {
	productService *domain.Service
}

// NewGetAllProductsUseCase creates a new instance of GetAllProductsUseCase
func NewGetAllProductsUseCase(productService *domain.Service) *GetAllProductsUseCase {
	return &GetAllProductsUseCase{
		productService: productService,
	}
}

// Execute runs the use case
func (uc *GetAllProductsUseCase) Execute(ctx context.Context) (*GetAllProductsOutput, error) {
	// Call domain service to get all products
	products, err := uc.productService.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}

	// Map domain entities to output
	output := &GetAllProductsOutput{
		Products: make([]ProductOutput, len(products)),
	}

	for i, p := range products {
		// Map categories
		categories := make([]CategoryOutput, 0, len(p.Categories()))
		for _, c := range p.Categories() {
			categories = append(categories, CategoryOutput{
				ID:   c.ID().String(),
				Name: c.Name().String(),
			})
		}

		output.Products[i] = ProductOutput{
			ID:          p.ID().String(),
			Name:        p.Name().String(),
			Description: p.Description().String(),
			Price:       p.Price().Amount(),
			Currency:    p.Price().Currency(),
			Stock:       p.Stock().Quantity(),
			Categories:  categories,
		}
	}

	return output, nil
}
