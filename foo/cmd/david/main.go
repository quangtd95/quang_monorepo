package main

import (
	"github.com/labstack/echo/v4"
	"github.com/quangtd95/quang_monorepo/foo/pkg"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return pkg.EchoResponse(c, http.StatusOK, pkg.SaySomething("david", "hello 9092"))
	})
	e.Logger.Fatal(e.Start(":9092"))
}
