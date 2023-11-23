package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Count struct {
	Count int `json:"count"`
}
var count *Count

func main() {
	count = &Count{
		Count: 0,
	}

	e := echo.New()
	e.GET("/count", GetCount)
	e.Logger.Fatal(e.Start(":8080"))
}

func GetCount(c echo.Context) error {
	return c.JSON(http.StatusOK, count)
}
