package main

import (
	"github.com/labstack/echo"
)

var secret = []byte("secret")

func main() {
	router := echo.New()
	router.Static("/", "views")
	router.GET("/status", status)
	router.GET("/get-token", getToken)

	dataset := router.Group("/dataset", jwtMiddleware, datasetLogger)
	dataset.GET("/employees", getEmployees)
	dataset.GET("/clients", getClients)

	router.Logger.Fatal(router.Start(":3300"))
}
