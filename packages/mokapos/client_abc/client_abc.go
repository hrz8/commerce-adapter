package main

import (
	"aiconec/commerce-adapter/core"
	"aiconec/commerce-adapter/pkg/proxy"
	"context"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	BASE_URL = "https://altalune.id/commerce"
)

var app *fiber.App
var fiberDoFunc *proxy.FiberProxy

func Main(ctx context.Context, event core.DigitalOceanParameters) (*core.DigitalOceanHTTPResponse, error) {
	fmt.Println(fmt.Sprintf("params: %+v\n", event))
	functionName := ctx.Value("function_name").(string)
	namespace := ctx.Value("namespace").(string)

	extractedPath := strings.TrimPrefix(functionName, "/"+namespace)
	ctx = context.WithValue(ctx, "app_host", BASE_URL+extractedPath)

	fmt.Println("ctx:", functionName, namespace, BASE_URL+extractedPath)
	fmt.Println(fmt.Sprintf("path 0 '%s'", event.HTTP.Path))

	app = fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World Fiber!")
	})

	fiberDoFunc = proxy.New(app)

	return fiberDoFunc.ProxyWithContext(ctx, event)

	// fmt.Println(fmt.Sprintf("params: %+v\n", event))
	// jsonString, err := json.Marshal(event)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println("JSON String:", string(jsonString))

	// host := ctx.Value("api_host").(string)
	// functionName := ctx.Value("function_name").(string)
	// namespace := ctx.Value("namespace").(string)

	// extractedPath := strings.TrimPrefix(functionName, "/"+namespace)
	// ctx = context.WithValue(ctx, "app_host", BASE_URL+extractedPath)

	// appHost := ctx.Value("app_host").(string)
	// namespace2 := ctx.Value("namespace").(string)

	// fmt.Println("ctx:", host, functionName, appHost, namespace2)
	// fmt.Println("cookie:", event.Headers["cookie"])

	// return &DigitalOceanHTTPResponse{
	// 	Body: fmt.Sprintf("Hello %s!", "stranger"),
	// }, nil
}
