package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(port())

	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//hello mwssage endpoint with a JSON response from map
	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello Gin framework"})
	})

	//get all books
	engine.GET("/api/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, AllBooks())
	})

	//Run server on PORT
	engine.Run(port())
}

func port() string {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
