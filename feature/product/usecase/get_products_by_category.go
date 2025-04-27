package product

import (
	"context"

	domain "sago-sample/feature/product/domain"
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
	productService *domain.Service
}

// NewGetProductsByCategoryUseCase creates a new instance of GetProductsByCategoryUseCase
func NewGetProductsByCategoryUseCase(productService *domain.Service) *GetProductsByCategoryUseCase {
	return &GetProductsByCategoryUseCase{
		productService: productService,
	}
}

// Execute runs the use case
func (uc *GetProductsByCategoryUseCase) Execute(ctx context.Context, input GetProductsByCategoryInput) (*GetProductsByCategoryOutput, error) {
	categoryID, err := domain.NewCategoryID(input.CategoryID)
	if err != nil {
		return nil, err
	}

	products, err := uc.productService.GetProductsByCategory(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	output := &GetProductsByCategoryOutput{
		Products: make([]ProductOutput, len(products)),
	}

	for i, p := range products {
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
