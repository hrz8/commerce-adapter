package main

import (
	App "aiconec/commerce-adapter/app"
	"aiconec/commerce-adapter/proxy"
	"context"

	"github.com/hrz8/do-function-go-proxy/core"

	"github.com/gofiber/fiber/v2"
)

const (
	FUNCTION_NAMESPACE = "commerce"
	BASE_URL           = "https://altalune.id/" + FUNCTION_NAMESPACE
)

var app *fiber.App
var adapter App.Adapter

func Main(_ctx context.Context, event core.DigitalOceanParameters) (*core.DigitalOceanHTTPResponse, error) {
	ctx := proxy.NewProxyContext(_ctx, BASE_URL)
	pCtx := ctx.Background()

	app = App.FiberApp(pCtx, FUNCTION_NAMESPACE)
	adapter = App.FiberAdapter(app)

	return adapter.ProxyWithContext(pCtx, event)
}
