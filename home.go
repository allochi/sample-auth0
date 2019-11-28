package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func home(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, world!")
}

func status(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Not Implemented")
}
