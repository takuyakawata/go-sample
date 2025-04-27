package product

import (
	"context"
	"errors"

	domain "sago-sample/feature/product/domain"
)

type GetProductInput struct {
	ID string
}

type GetProductOutput struct {
	ID          string
	Name        string
	Description string
	Price       uint
	Currency    string
	Stock       uint
	Categories  []CategoryOutput
}

type GetProductUseCase struct {
	repo domain.Repository
}

func NewGetProductUseCase(repo domain.Repository) *GetProductUseCase {
	return &GetProductUseCase{
		repo: repo,
	}
}

// Execute runs the use case
func (uc *GetProductUseCase) Execute(ctx context.Context, input GetProductInput) (*GetProductOutput, error) {
	// Create value object
	productID, err := domain.NewProductID(input.ID)
	if err != nil {
		return nil, err
	}

	// Call domain service to get product
	foundProduct, err := uc.repo.FindByID(ctx, productID)
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

	return &GetProductOutput{
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
	repo domain.Repository
}

// NewGetAllProductsUseCase creates a new instance of GetAllProductsUseCase
func NewGetAllProductsUseCase(repo domain.Repository) *GetAllProductsUseCase {
	return &GetAllProductsUseCase{repo: repo}
}

// Execute runs the use case
func (uc *GetAllProductsUseCase) Execute(ctx context.Context) (*GetAllProductsOutput, error) {
	products, err := uc.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	output := &GetAllProductsOutput{
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
