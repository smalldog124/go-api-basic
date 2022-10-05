package main

import (
	"basic/customer"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

func main() {
	url := "localhost:27017"
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	customerAPI := customer.Handlers{
		CustomerRepo: customer.RepositoryMongo{
			DBName:       "buddy",
			DBConnection: session,
		},
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/api/v1/customer", customerAPI.CreatedCustomerHandlers)
	r.Run(":8000")
}
