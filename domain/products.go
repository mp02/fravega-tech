package domain

import (
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

// ProductRepository define las operaciones del repositorio
type ProductRepository interface {
	Create(product *Product) (Product, error)
	GetAll(page, limit int) ([]Product, error)
	Delete(id string) (Product, error)
	Update(id string, updateProduct *UpdateProduct) (Product, error)
	GetProductsByFilters(filters ProductFilters, page int, limit int) ([]Product, error)
}
