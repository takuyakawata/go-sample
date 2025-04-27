package product

import (
	"context"
	"errors"

	domain "sago-sample/feature/product/domain"
)

type RemoveCategoryFromProductInput struct {
	ProductID  string
	CategoryID string
}

type RemoveCategoryFromProductOutput struct {
	ProductID   string
	Name        string
	Description string
	Price       uint
	Currency    string
	Stock       uint
	Categories  []CategoryOutput
}

type RemoveCategoryFromProductUseCase struct {
	productService *domain.Service
}

func NewRemoveCategoryFromProductUseCase(productService *domain.Service) *RemoveCategoryFromProductUseCase {
	return &RemoveCategoryFromProductUseCase{
		productService: productService,
	}
}

func (uc *RemoveCategoryFromProductUseCase) Execute(ctx context.Context, input RemoveCategoryFromProductInput) (*RemoveCategoryFromProductOutput, error) {
	productID, err := domain.NewProductID(input.ProductID)
	if err != nil {
		return nil, err
	}

	categoryID, err := domain.NewCategoryID(input.CategoryID)
	if err != nil {
		return nil, err
	}

	updatedProduct, err := uc.productService.RemoveCategoryFromProduct(ctx, productID, categoryID)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	categories := make([]CategoryOutput, 0, len(updatedProduct.Categories()))
	for _, c := range updatedProduct.Categories() {
		categories = append(categories, CategoryOutput{
			ID:   c.ID().String(),
			Name: c.Name().String(),
		})
	}

	return &RemoveCategoryFromProductOutput{
		ProductID:   updatedProduct.ID().String(),
		Name:        updatedProduct.Name().String(),
		Description: updatedProduct.Description().String(),
		Price:       updatedProduct.Price().Amount(),
		Currency:    updatedProduct.Price().Currency(),
		Stock:       updatedProduct.Stock().Quantity(),
		Categories:  categories,
	}, nil
}
