package main

import (
	docore "aiconec/commerce-adapter/core/do"
	"context"
	"fmt"
)

func Main(ctx context.Context, req docore.DigitalOceanHTTPRequest) (*docore.DigitalOceanHTTPResponse, error) {
	return &docore.DigitalOceanHTTPResponse{
		Body: fmt.Sprintf("Hello %s!", req.QueryString),
	}, nil
}
