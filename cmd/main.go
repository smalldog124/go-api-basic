package main

import (
	"basic/customer"
	"basic/product"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/jmoiron/sqlx"
)

func main() {
	url := "localhost:27017"
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	db, err := sqlx.Connect("postgres", "user=admin password=example dbname=liza host=localhost port=5433 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	customerAPI := customer.Handlers{
		CustomerRepo: customer.RepositoryMongo{
			DBName:       "buddy",
			DBConnection: session,
		},
	}
	productAPI := product.Handlers{
		ProductRepo: product.RepositoryPg{
			DBConnection: db,
		},
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/api/v1/customer", customerAPI.CreatedCustomerHandlers)
	r.POST("/api/v1/product", productAPI.CreatedProductHandlers)
	r.Run(":8000")
}
