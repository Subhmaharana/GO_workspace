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
