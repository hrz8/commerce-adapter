package main

import (
	App "aiconec/commerce-adapter/app"
	"context"

	"github.com/hrz8/do-function-go-proxy/core"
	"github.com/hrz8/do-function-go-proxy/pkg/proxy"

	"github.com/gofiber/fiber/v2"
)

const (
	FUNCTION_NAMESPACE = "commerce"
	BASE_URL           = "https://altalune.id/" + FUNCTION_NAMESPACE
)

var app *fiber.App
var adapter App.Adapter

func Main(ctx context.Context, event core.DigitalOceanParameters) (*core.DigitalOceanHTTPResponse, error) {
	pCtx := proxy.NewContext(ctx, BASE_URL).Background()

	app = App.FiberApp(pCtx, FUNCTION_NAMESPACE)
	adapter = App.FiberAdapter(app)

	return adapter.ProxyWithContext(pCtx, event)
}
