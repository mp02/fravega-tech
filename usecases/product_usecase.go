package usecases

import (
	"errors"
	"time"

	"github.com/mp02/fravega-tech/domain"
)

type ProductUseCase struct {
	Repo domain.ProductRepository
}

func NewProductUseCase(repo domain.ProductRepository) *ProductUseCase {
	return &ProductUseCase{Repo: repo}
}

// CreateProduct crea un nuevo producto
func (uc *ProductUseCase) CreateProduct(product *domain.Product) (domain.Product, error) {
	if product.Name == "" || product.Price <= 0 {
		return domain.Product{}, errors.New("invalid product data")
	}
	product.CreatedAt = time.Now()
	return uc.Repo.Create(product)
}

// GetAllProducts devuelve todos los productos
func (uc *ProductUseCase) GetAllActiveProducts() ([]domain.Product, error) {
	return uc.Repo.GetAll()
}

// GetAllProducts devuelve todos los productos
func (uc *ProductUseCase) DeleteProduct(id string) (domain.Product, error) {
	return uc.Repo.Delete(id)
}

func (uc *ProductUseCase) UpdateProductByID(id string, product *domain.UpdateProduct) (domain.Product, error) {
	return uc.Repo.Update(id, product)
}

func (u *ProductUseCase) GetProductsWithFilters(filters domain.ProductFilters) ([]domain.Product, error) {
	return u.Repo.GetProductsByFilters(filters)
}
