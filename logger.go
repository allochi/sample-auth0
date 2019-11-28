package main

import (
	"log"

	"github.com/labstack/echo"
)

func datasetLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("We are in the middleware!!!")
		return next(c)
	}
}
