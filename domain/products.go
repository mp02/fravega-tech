package domain

import (
	"errors"
	"time"
)

type Product struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Name        string    `bson:"name" json:"name"`
	Description string    `bson:"description" json:"description"`
	Price       float64   `bson:"price" json:"price"`
	Categories  []string  `bson:"categories" json:"categories"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
	IsDeleted   bool      `bson:"is_deleted" json:"is_deleted,omitempty"`
	ImagesURL   []string  `bson:"images_url" json:"images_url"`
}

type UpdateProduct struct {
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	Price       *float64  `json:"price,omitempty"`
	Categories  *[]string `json:"categories,omitempty"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
	IsDeleted   *bool     `json:"is_deleted,omitempty"`
}

type ProductFilters struct {
	Name       *string   `form:"name"`
	MinPrice   *float64  `form:"min_price"`
	MaxPrice   *float64  `form:"max_price"`
	Categories *[]string `form:"categories"`
	IsDeleted  *bool     `form:"is_deleted"`
}

// ProductRepository define las operaciones del repositorio
type ProductRepository interface {
	Create(product *Product) (Product, error)
	GetAll() ([]Product, error)
	Delete(id string) (Product, error)
	Update(id string, updateProduct *UpdateProduct) (Product, error)
	GetProductsByFilters(filters ProductFilters) ([]Product, error)
}

func (f *ProductFilters) Validate() error {
	if *f.MinPrice > *f.MaxPrice && *f.MaxPrice > 0 {
		return errors.New("min_price cannot be greater than max_price")
	}
	return nil
}

func (f *ProductFilters) AreFiltersEmpty() bool {
	if f.MinPrice == nil && f.MaxPrice == nil && f.Categories == nil && f.IsDeleted == nil {
		return true
	}
	return false
}
