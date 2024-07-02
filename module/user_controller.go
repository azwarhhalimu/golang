package module

import (
	"context"
	"golang/prisma/db"

	"github.com/labstack/echo/v4"
)

func Okos(res echo.Context) error {
	ctx := context.Background()
	client := db.NewClient()
	client.Prisma.Connect()
	hasil, _ := client.Kecamatan.FindMany(
		db.Kecamatan.IDKecamatan.Contains("01"),
	).With(
		db.Kecamatan.Kelurahan.Fetch().With(
			db.Kelurahan.BukuTanah.Fetch(),
			db.Kelurahan.Kecamatan.Fetch(),
		),
	).Exec(ctx)
	return res.JSON(200, map[string]interface{}{
		"status":  "ok",
		"message": "saya pergi kep asar",
		"data":    hasil,
		"okos":    hasil[0].Kelurahan()[0].Kelurahan,
	})
}
