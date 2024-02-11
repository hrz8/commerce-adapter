package app

import (
	"context"

	"github.com/hrz8/do-function-go-proxy/core"
)

type Adapter interface {
	ProxyWithContext(ctx context.Context, params core.DigitalOceanParameters) (*core.DigitalOceanHTTPResponse, error)
}
