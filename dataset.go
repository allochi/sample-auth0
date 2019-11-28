package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Employee struct {
	Name  string
	Title string
}

func getEmployees(ctx echo.Context) error {
	// Assume this is database call
	employees := []Employee{
		{"Ali Anwar", "typer"},
		{"Safae", "The smartie"},
		{"Ariel", "The data guy"},
	}

	return ctx.JSON(http.StatusOK, employees)
}

func getClients(ctx echo.Context) error {
	// Assume this is database call
	clients := []Employee{
		{"Tom Cruse", "Nah!"},
	}

	return ctx.JSON(http.StatusOK, clients)
}
