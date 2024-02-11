package main

import (
	"aiconec/commerce-adapter/config"
	App "aiconec/commerce-adapter/internal/app"
	"context"

	"github.com/hrz8/do-function-go-proxy/core"
	echoadapter "github.com/hrz8/do-function-go-proxy/pkg/adapter/echo"
	"github.com/hrz8/do-function-go-proxy/pkg/proxy"
	"github.com/labstack/echo/v4"
)

var echoApp *echo.Echo

func Main(ctx context.Context, event core.DigitalOceanParameters) (*core.DigitalOceanHTTPResponse, error) {
	appConfig := config.New()

	pCtx := proxy.NewContext(ctx, appConfig.GetBaseURL()).Background()

	echoApp = echo.New()

	app := App.New(appConfig, echoApp)
	app.Load(pCtx)

	return echoadapter.New(echoApp).ProxyWithContext(pCtx, event)
}
