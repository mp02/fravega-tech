package interfaces

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mp02/fravega-tech/domain"
	"github.com/mp02/fravega-tech/usecases"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductHandler struct {
	UseCase *usecases.ProductUseCase
}

func NewProductHandler(uc *usecases.ProductUseCase) *ProductHandler {
	return &ProductHandler{UseCase: uc}
}

// CreateProduct godoc
// @Summary Crea un nuevo producto
// @Description Permite crear un nuevo producto con la información proporcionada
// @Tags productos
// @ID create-product
// @Accept json
// @Produce json
// @Param product body domain.Product true "Producto a crear"
// @Success 201 {object} domain.Product
// @Failure 400 {object} error
// @Router /products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	newProduct, err := h.UseCase.CreateProduct(&product)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newProduct)
}

// GetProducts godoc
// @Summary Obtiene una lista de productos
// @Description Devuelve todos los productos, con posibilidad de aplicar filtros
// @Tags productos
// @ID get-products
// @Produce json
// @Param categories query []string false "Filtros por categorías"
// @Param min_price query float64 false "Filtro por precio mínimo"
// @Param max_price query float64 false "Filtro por precio máximo"
// @Success 200 {array} domain.Product
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	var filters domain.ProductFilters
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filters"})
		return
	}
	if filters.AreFiltersEmpty() {
		products, err := h.UseCase.GetAllActiveProducts()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching all active products"})
			return
		}
		c.JSON(http.StatusOK, products)
		return
	}

	products, err := h.UseCase.GetProductsWithFilters(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching filtered products"})
		return
	}

	c.JSON(http.StatusOK, products)
	return
}

// DeleteProduct godoc
// @Summary Elimina un producto por su ID
// @Description Marca como eliminado un producto específico
// @Tags productos
// @ID delete-product
// @Param id path string true "Producto ID"
// @Success 200 {object} domain.Product
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	productID := c.Param("id") // Get the product ID from the URL

	// Validate if the product ID is a valid ObjectID
	_, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	products, err := h.UseCase.DeleteProduct(productID)
	if err != nil && products.ID == "" {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "product:" + productID + " not found"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// UpdateProductByID godoc
// @Summary Actualiza un producto por su ID
// @Description Permite actualizar la información de un producto específico
// @Tags productos
// @ID update-product
// @Accept json
// @Produce json
// @Param id path string true "Producto ID"
// @Param product body domain.UpdateProduct true "Producto actualizado"
// @Success 200 {object} domain.Product
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /products/{id} [patch]
func (h *ProductHandler) UpdateProductByID(c *gin.Context) {
	productID := c.Param("id") // Get the product ID from the URL

	// Validate if the product ID is a valid ObjectID
	_, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	var updateProductJSON domain.UpdateProduct
	if err := c.ShouldBindJSON(&updateProductJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	products, err := h.UseCase.UpdateProductByID(productID, &updateProductJSON)
	if err != nil && products.ID == "" {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "product:" + productID + " not found"})
		return
	}

	c.JSON(http.StatusOK, products)
}
