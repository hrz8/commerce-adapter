package app

import (
	"aiconec/commerce-adapter/internal/port"
	"context"
	"fmt"

	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/labstack/echo/v4"
)

type AppRegistry struct {
	cfg         port.ServiceConfig
	echoApp     *echo.Echo
	usecases    *hashmap.Map
	controllers *hashmap.Map
}

func New(cfg port.ServiceConfig, echoApp *echo.Echo) *AppRegistry {
	return &AppRegistry{
		cfg:     cfg,
		echoApp: echoApp,

		usecases:    nil,
		controllers: nil,
	}
}

func (r *AppRegistry) Load(ctx context.Context) {
	path := ctx.Value("trailing_path").(string)

	r.usecases = r.loadUsecases()
	r.controllers = r.loadControllers()

	basePath := r.echoApp.Group(fmt.Sprintf("/%s%s", r.cfg.GetFunctionName(), path))

	for _, key := range r.controllers.Keys() {
		_ctrl, exist := r.controllers.Get(key)
		ctrl, ok := _ctrl.(port.Controller)
		if exist && ok {
			ctrl.Routers(basePath)
		}
	}
}
