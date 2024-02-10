package main

import (
	"aiconec/commerce-adapter/core"
	echoadapter "aiconec/commerce-adapter/pkg/adapter/echo"
	fiberadapter "aiconec/commerce-adapter/pkg/adapter/fiber"
	"context"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	FUNCTION_NAMESPACE = "commerce"
	BASE_URL           = "https://altalune.id/" + FUNCTION_NAMESPACE
)

var app *echo.Echo
var adapter Adapter

type Adapter interface {
	ProxyWithContext(ctx context.Context, params core.DigitalOceanParameters) (*core.DigitalOceanHTTPResponse, error)
}

func initCtx(ctx context.Context) context.Context {
	// /ap-xxx-xxx/mokapos/client_abc
	functionName := ctx.Value("function_name").(string)
	// ap-xxx-xxx
	namespace := ctx.Value("namespace").(string)
	// /mokapos/client_abc
	extractedPath := strings.TrimPrefix(functionName, "/"+namespace)

	ctx = context.WithValue(ctx, "trailing_path", extractedPath)
	ctx = context.WithValue(ctx, "app_host", BASE_URL+extractedPath)

	return ctx
}

func fiberApp(ctx context.Context) *fiber.App {
	path := ctx.Value("trailing_path").(string)

	app := fiber.New()
	router := app.Group(fmt.Sprintf("/%s%s", FUNCTION_NAMESPACE, path))

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World Fiber!")
	})

	router.Get("/uhuy", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World Huy!")
	})

	return app
}

func fiberAdapter(app *fiber.App) Adapter {
	return fiberadapter.New(app)
}

func echoApp(ctx context.Context) *echo.Echo {
	path := ctx.Value("trailing_path").(string)

	e := echo.New()
	e.Pre(middleware.AddTrailingSlash())

	router := e.Group(fmt.Sprintf("/%s%s", FUNCTION_NAMESPACE, path))

	router.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World Echo!")
	})

	router.GET("/uhuy", func(c echo.Context) error {
		return c.String(200, "Hello, World Echuy!")
	})

	return e
}

func echoAdapter(app *echo.Echo) Adapter {
	return echoadapter.New(app)
}

func Main(ctx context.Context, event core.DigitalOceanParameters) (*core.DigitalOceanHTTPResponse, error) {
	ctx = initCtx(ctx)
	app = echoApp(ctx)
	adapter = echoAdapter(app)

	return adapter.ProxyWithContext(ctx, event)
}
