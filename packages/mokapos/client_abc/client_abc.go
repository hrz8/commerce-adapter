package main

import (
	"aiconec/commerce-adapter/core"
	fiberadapter "aiconec/commerce-adapter/pkg/adapter/fiber"
	"context"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	FUNCTION_NAMESPACE = "commerce"
	FUNCTION_PACKAGE   = "mokapos"
	FUNCTION_NAME      = "client_abc"
	BASE_URL           = "https://altalune.id/" + FUNCTION_NAMESPACE
)

var app *fiber.App
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
	// baseFunction := app.Group(fmt.Sprintf("/%s/%s/%s", FUNCTION_NAME, FUNCTION_PACKAGE, FUNCTION_NAME))

	app.Get("/commerce/mokapos/client_abc", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World Fiber!")
	})

	app.Get("/commerce/mokapos/client_abc/uhuy", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World Huy!")
	})

	return app
}

func fiberAdapter(app *fiber.App) Adapter {
	return fiberadapter.New(app)
}

func Main(ctx context.Context, event core.DigitalOceanParameters) (*core.DigitalOceanHTTPResponse, error) {
	app = fiberApp()
	adapter = fiberAdapter(app)

	return adapter.ProxyWithContext(DoCtx(ctx), event)
}
