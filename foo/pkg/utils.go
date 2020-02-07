package pkg

import "github.com/labstack/echo/v4"

func SaySomething(servicename string, text string) string {
	return servicename + " said that: " + text
}

func EchoResponse(ctx echo.Context, code int, message string) error {
	return ctx.String(code, message)
}
