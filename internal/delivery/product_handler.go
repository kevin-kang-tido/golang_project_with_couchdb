package delivery

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"golang_with_couchdb2/internal/domain/entities"
	"golang_with_couchdb2/internal/usecases/interactors"
)

type ProductHandler struct {
	ProductInteractor *interactors.ProductInteractor
}

func NewProductHandler(interactor *interactors.ProductInteractor) *ProductHandler {
	return &ProductHandler{ProductInteractor: interactor}
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
    products, err := h.ProductInteractor.GetAllProducts()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product entities.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.ProductInteractor.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}



func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id := c.Param("_id")
	product, err := h.ProductInteractor.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) UpdateProductByID(c *gin.Context) {
    id := c.Param("_id") // Get the ID from the URL
    var product entities.Product

    // Bind the JSON body to the product struct
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Set the product ID
    product.ID = id // Set the ID from the URL to the product object

    // Call the interactor to update the product
    err := h.ProductInteractor.UpdateProduct(&product)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Provide detailed error message
        return
    }

    // Respond with the updated product
    c.JSON(http.StatusOK, product)
}


func (h *ProductHandler) DeleteProductByID(c *gin.Context) {
	id := c.Param("_id")
	if err := h.ProductInteractor.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
