package app

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	fiberadapter "github.com/hrz8/do-function-go-proxy/pkg/adapter/fiber"
)

func FiberApp(ctx context.Context, functionNamespace string) *fiber.App {
	path := ctx.Value("trailing_path").(string)

	app := fiber.New()
	router := app.Group(fmt.Sprintf("/%s%s", functionNamespace, path))

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World Fiber!")
	})

	router.Get("/uhuy", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World Huy!")
	})

	router.Get("/owo/iwi", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World Fibwo!")
	})

	return app
}

func FiberAdapter(app *fiber.App) Adapter {
	return fiberadapter.New(app)
}
