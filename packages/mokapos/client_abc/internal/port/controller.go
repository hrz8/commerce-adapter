package port

import "github.com/labstack/echo/v4"

type Controller interface {
	Routers(r *echo.Group)
}
