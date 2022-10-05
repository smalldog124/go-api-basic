package customer

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	CustomerRepo CustomerRepo
}

type CustomerRepo interface {
	CreatedCustomer(customer CustomerInfo) error
}

func (h Handlers) CreatedCustomerHandlers(context *gin.Context) {
	var requestBody CustomerInfo
	err := context.ShouldBindJSON(&requestBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err = h.CustomerRepo.CreatedCustomer(requestBody)
	if err != nil {
		log.Println("CustomerRepo.CreatedCustomer", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "success"})
}
