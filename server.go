package main

import (
	"golang/module"

	"github.com/labstack/echo/v4"
)

func main() {
	route := echo.New()

	route.GET("/", index)
	route.GET("/azwar", module.Okos)
	route.GET("/halaman2", module.Azwar)

	route.Logger.Fatal(route.Start(":2500"))
}
func index(res echo.Context) error {
	return res.String(200, "halaman index")
}
func halaman2(res echo.Context) error {
	return res.String(200, "halao semuanya")
}
