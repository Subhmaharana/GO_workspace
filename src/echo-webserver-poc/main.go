package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	fmt.Printf("Starting Echo Web server...")

	e := echo.New()
	e.GET("/", helloFunc)
	e.Start(":8081")

}

func helloFunc(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from the Web Server!!")
}
