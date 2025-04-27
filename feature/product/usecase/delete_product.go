package product

import (
	"context"
	"errors"

	domain "sago-sample/feature/product/domain"
)

// DeleteProductInput represents the input data for deleting a product
type DeleteProductInput struct {
	ID string
}

// DeleteProductUseCase defines the use case for deleting a product
type DeleteProductUseCase struct {
	productService *domain.Service
}

// NewDeleteProductUseCase creates a new instance of DeleteProductUseCase
func NewDeleteProductUseCase(productService *domain.Service) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		productService: productService,
	}
}

// Execute runs the use case
func (uc *DeleteProductUseCase) Execute(ctx context.Context, input DeleteProductInput) error {
	productID, err := domain.NewProductID(input.ID)
	if err != nil {
		return err
	}

	err = uc.productService.DeleteProduct(ctx, productID)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			return errors.New("product not found")
		}
		return err
	}

	return nil
}
