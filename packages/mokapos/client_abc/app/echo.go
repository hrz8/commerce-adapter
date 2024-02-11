package app

import (
	"context"
	"fmt"

	echoadapter "github.com/hrz8/do-function-go-proxy/pkg/adapter/echo"
	"github.com/labstack/echo/v4"
)

func EchoApp(ctx context.Context, functionNamespace string) *echo.Echo {
	path := ctx.Value("trailing_path").(string)

	e := echo.New()

	router := e.Group(fmt.Sprintf("/%s%s", functionNamespace, path))

	router.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World Echo!")
	})

	router.GET("/uhuy", func(c echo.Context) error {
		return c.String(200, "Hello, World Echuy!")
	})

	router.GET("/owo/iwi", func(c echo.Context) error {
		return c.String(200, "Hello, World Echuy!")
	})

	return e
}

func EchoAdapter(app *echo.Echo) Adapter {
	return echoadapter.New(app)
}
