package product

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// ProductID represents the unique identifier for a product
type ProductID string

// NewProductID creates a new ProductID
func NewProductID(id string) (ProductID, error) {
	if strings.TrimSpace(id) == "" {
		return "", errors.New("product id cannot be empty")
	}
	return ProductID(id), nil
}

// String returns the string representation of the ProductID
func (id ProductID) String() string {
	return string(id)
}

// IsEmpty checks if the ProductID is empty
func (id ProductID) IsEmpty() bool {
	return strings.TrimSpace(string(id)) == ""
}

// MustNewProductID creates a new ProductID and panics if validation fails
func MustNewProductID(id string) ProductID {
	productID, err := NewProductID(id)
	if err != nil {
		panic(err)
	}
	return productID
}

// ProductName represents the name of a product
type ProductName string

// NewProductName creates a new ProductName
func NewProductName(name string) (ProductName, error) {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return "", errors.New("product name cannot be empty")
	}
	if len(trimmedName) > 100 {
		return "", errors.New("product name cannot exceed 100 characters")
	}
	return ProductName(trimmedName), nil
}

// String returns the string representation of the ProductName
func (n ProductName) String() string {
	return string(n)
}

// IsEmpty checks if the ProductName is empty
func (n ProductName) IsEmpty() bool {
	return strings.TrimSpace(string(n)) == ""
}

// MustNewProductName creates a new ProductName and panics if validation fails
func MustNewProductName(name string) ProductName {
	productName, err := NewProductName(name)
	if err != nil {
		panic(err)
	}
	return productName
}

// ProductDescription represents the description of a product
type ProductDescription string

// NewProductDescription creates a new ProductDescription
func NewProductDescription(description string) (ProductDescription, error) {
	trimmedDesc := strings.TrimSpace(description)
	if len(trimmedDesc) > 1000 {
		return "", errors.New("product description cannot exceed 1000 characters")
	}
	return ProductDescription(trimmedDesc), nil
}

// String returns the string representation of the ProductDescription
func (d ProductDescription) String() string {
	return string(d)
}

// MustNewProductDescription creates a new ProductDescription and panics if validation fails
func MustNewProductDescription(description string) ProductDescription {
	productDescription, err := NewProductDescription(description)
	if err != nil {
		panic(err)
	}
	return productDescription
}

// Price represents the monetary value of a product
type Price struct {
	amount   uint
	currency string
}

// NewPrice creates a new Price
func NewPrice(amount uint, currency string) (Price, error) {
	if amount == 0 {
		return Price{}, errors.New("price amount cannot be zero")
	}

	currency = strings.ToUpper(strings.TrimSpace(currency))
	if currency == "" {
		return Price{}, errors.New("currency cannot be empty")
	}

	// Simple currency code validation (3 uppercase letters)
	match, _ := regexp.MatchString("^[A-Z]{3}$", currency)
	if !match {
		return Price{}, errors.New("invalid currency format, must be 3 uppercase letters")
	}

	return Price{
		amount:   amount,
		currency: currency,
	}, nil
}

// Amount returns the amount of the price
func (p Price) Amount() uint {
	return p.amount
}

// Currency returns the currency of the price
func (p Price) Currency() string {
	return p.currency
}

// String returns the string representation of the Price
func (p Price) String() string {
	return fmt.Sprintf("%d %s", p.amount, p.currency)
}

// MustNewPrice creates a new Price and panics if validation fails
func MustNewPrice(amount uint, currency string) Price {
	price, err := NewPrice(amount, currency)
	if err != nil {
		panic(err)
	}
	return price
}

// Stock represents the available quantity of a product
type Stock struct {
	quantity uint
}

// NewStock creates a new Stock
func NewStock(quantity uint) Stock {
	return Stock{quantity: quantity}
}

// Quantity returns the quantity of the stock
func (s Stock) Quantity() uint {
	return s.quantity
}

// Increase increases the stock quantity by the given amount
func (s *Stock) Increase(amount uint) {
	s.quantity += amount
}

// Decrease decreases the stock quantity by the given amount
func (s *Stock) Decrease(amount uint) error {
	if amount > s.quantity {
		return errors.New("insufficient stock")
	}
	s.quantity -= amount
	return nil
}

// IsAvailable checks if the product is available (has stock)
func (s Stock) IsAvailable() bool {
	return s.quantity > 0
}

// String returns the string representation of the Stock
func (s Stock) String() string {
	return fmt.Sprintf("%d units", s.quantity)
}
