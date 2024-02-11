package proxy

import (
	"context"
	"strings"
)

type Context struct {
	ctx     context.Context
	baseURL string
}

func NewProxyContext(ctx context.Context, baseURL string) *Context {
	return &Context{ctx, baseURL}
}

func (pc *Context) Background() context.Context {
	// /ap-xxx-xxx/mokapos/client_abc
	functionName := pc.ctx.Value("function_name").(string)
	// ap-xxx-xxx
	namespace := pc.ctx.Value("namespace").(string)
	// /mokapos/client_abc
	extractedPath := strings.TrimPrefix(functionName, "/"+namespace)

	pc.ctx = context.WithValue(pc.ctx, "trailing_path", extractedPath)
	pc.ctx = context.WithValue(pc.ctx, "app_host", pc.baseURL+extractedPath)

	return pc.ctx
}
