package main

import (
	"context"
	"fmt"
)

type DigitalOceanHTTPRequest struct {
	Body            string              `json:"body"`
	Headers         map[string][]string `json:"headers"`
	IsBase64Encoded bool                `json:"isBase64Encoded"`
	Method          string              `json:"method"`
	Path            string              `json:"path"`
	QueryString     string              `json:"queryString"`
}

type DigitalOceanHTTPResponse struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

func Main(ctx context.Context, req DigitalOceanHTTPRequest) (*DigitalOceanHTTPResponse, error) {
	fmt.Println(ctx)
	fmt.Println(req)
	return &DigitalOceanHTTPResponse{
		Body: fmt.Sprintf("Hello %s!", req.QueryString),
	}, nil
}
