package product

import (
	"errors"
	"strings"
)

// CategoryID represents the unique identifier for a category
type CategoryID string

// NewCategoryID creates a new CategoryID
func NewCategoryID(id string) (CategoryID, error) {
	if strings.TrimSpace(id) == "" {
		return "", errors.New("category id cannot be empty")
	}
	return CategoryID(id), nil
}

// String returns the string representation of the CategoryID
func (id CategoryID) String() string {
	return string(id)
}

// IsEmpty checks if the CategoryID is empty
func (id CategoryID) IsEmpty() bool {
	return strings.TrimSpace(string(id)) == ""
}

// CategoryName represents the name of a category
type CategoryName string

// NewCategoryName creates a new CategoryName
func NewCategoryName(name string) (CategoryName, error) {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return "", errors.New("category name cannot be empty")
	}
	if len(trimmedName) > 50 {
		return "", errors.New("category name cannot exceed 50 characters")
	}
	return CategoryName(trimmedName), nil
}

// String returns the string representation of the CategoryName
func (n CategoryName) String() string {
	return string(n)
}

// IsEmpty checks if the CategoryName is empty
func (n CategoryName) IsEmpty() bool {
	return strings.TrimSpace(string(n)) == ""
}

// Category represents a product category
type Category struct {
	id   CategoryID
	name CategoryName
}

// NewCategory creates a new Category
func NewCategory(id CategoryID, name CategoryName) (*Category, error) {
	if id.IsEmpty() {
		return nil, errors.New("category id cannot be empty")
	}
	if name.IsEmpty() {
		return nil, errors.New("category name cannot be empty")
	}
	return &Category{
		id:   id,
		name: name,
	}, nil
}

// ID returns the category's ID
func (c *Category) ID() CategoryID {
	return c.id
}

// Name returns the category's name
func (c *Category) Name() CategoryName {
	return c.name
}

// UpdateName updates the category's name
func (c *Category) UpdateName(name CategoryName) error {
	if name.IsEmpty() {
		return errors.New("category name cannot be empty")
	}
	c.name = name
	return nil
}
