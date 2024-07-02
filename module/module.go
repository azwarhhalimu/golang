package module

import "github.com/labstack/echo/v4"

func Azwar(res echo.Context) error {
	return res.String(200, "Azwar genteng")
}
func User(res echo.Context) error {
	return res.String(200, "halaman azwar buton")
}
