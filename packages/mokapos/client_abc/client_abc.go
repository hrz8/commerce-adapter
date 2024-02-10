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
)

const (
	FUNCTION_NAMESPACE = "commerce"
	FUNCTION_PACKAGE   = "mokapos"
	FUNCTION_NAME      = "client_abc"
	BASE_URL           = "https://altalune.id/" + FUNCTION_NAMESPACE
)

var app *echo.Echo
var adapter Adapter

type Adapter interface {
	ProxyWithContext(ctx context.Context, params core.DigitalOceanParameters) (*core.DigitalOceanHTTPResponse, error)
}

func DoCtx(ctx context.Context) context.Context {
	functionName := ctx.Value("function_name").(string)
	namespace := ctx.Value("namespace").(string)

	extractedPath := strings.TrimPrefix(functionName, "/"+namespace)

	return context.WithValue(ctx, "app_host", BASE_URL+extractedPath)
}

func fiberApp() *fiber.App {
	app := fiber.New()
	router := app.Group(fmt.Sprintf("/%s/%s/%s", FUNCTION_NAMESPACE, FUNCTION_PACKAGE, FUNCTION_NAME))

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

func echoApp() *echo.Echo {
	e := echo.New()
	router := e.Group(fmt.Sprintf("/%s/%s/%s", FUNCTION_NAMESPACE, FUNCTION_PACKAGE, FUNCTION_NAME))

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
	app = echoApp()
	adapter = echoAdapter(app)

	return adapter.ProxyWithContext(DoCtx(ctx), event)
}
