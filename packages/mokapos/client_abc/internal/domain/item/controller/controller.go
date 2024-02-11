package controller

import (
	"aiconec/commerce-adapter/internal/port"

	"github.com/labstack/echo/v4"
)

type JSONResp struct {
	Message string `json:"message"`
}

type ItemController struct {
	cfg     port.ServiceConfig
	usecase port.ServiceUsecase
}

func New(cfg port.ServiceConfig, usecase port.ServiceUsecase) *ItemController {
	return &ItemController{
		cfg,
		usecase,
	}
}

func (c *ItemController) Routers(r *echo.Group) {
	r.GET("/items", c.getItems)
	r.GET("/items/json", c.getItemsJSON)
	r.POST("/items/post", c.getItemsPOST)
}

func (c *ItemController) getItems(ctx echo.Context) error {
	result, err := c.usecase.GetItems(ctx.Request().Context())
	if err != nil {
		return err
	}

	return ctx.String(200, result)
}

func (c *ItemController) getItemsJSON(ctx echo.Context) error {
	result, err := c.usecase.GetItems(ctx.Request().Context())
	if err != nil {
		return err
	}

	return ctx.JSON(200, &JSONResp{
		Message: result,
	})
}

func (c *ItemController) getItemsPOST(ctx echo.Context) error {
	result, err := c.usecase.GetItems(ctx.Request().Context())
	if err != nil {
		return err
	}

	return ctx.JSON(200, &JSONResp{
		Message: result + " post",
	})
}
