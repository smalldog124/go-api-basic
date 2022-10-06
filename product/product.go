package product

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	ProductRepo ProductRepo
}

type ProductRepo interface {
	CreatedProduct(product Product) error
}

func (h Handlers) CreatedProductHandlers(context *gin.Context) {
	var requestBody Product
	err := context.ShouldBindJSON(&requestBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err = h.ProductRepo.CreatedProduct(requestBody)
	if err != nil {
		log.Println("ProductRepo.CreatedProduct", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "success"})
}
