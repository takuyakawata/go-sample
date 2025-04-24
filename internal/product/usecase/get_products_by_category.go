package product

import (
	"context"

	domain "sago-sample/internal/product/domain"
)

// GetProductsByCategoryInput represents the input data for getting products by category
type GetProductsByCategoryInput struct {
	CategoryID string
}

// GetProductsByCategoryOutput represents the output data after getting products by category
type GetProductsByCategoryOutput struct {
	Products []ProductOutput
}

// GetProductsByCategoryUseCase defines the use case for getting products by category
type GetProductsByCategoryUseCase struct {
	productService *product.Service
}

// NewGetProductsByCategoryUseCase creates a new instance of GetProductsByCategoryUseCase
func NewGetProductsByCategoryUseCase(productService *product.Service) *GetProductsByCategoryUseCase {
	return &GetProductsByCategoryUseCase{
		productService: productService,
	}
}

// Execute runs the use case
func (uc *GetProductsByCategoryUseCase) Execute(ctx context.Context, input GetProductsByCategoryInput) (*GetProductsByCategoryOutput, error) {
	// Create value object
	categoryID, err := product.NewCategoryID(input.CategoryID)
	if err != nil {
		return nil, err
	}

	// Call domain service to get products by category
	products, err := uc.productService.GetProductsByCategory(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	// Map domain entities to output
	output := &GetProductsByCategoryOutput{
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
