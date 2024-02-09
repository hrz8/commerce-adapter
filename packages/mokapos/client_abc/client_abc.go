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

type DigitalOceanParameters struct {
	HTTP DigitalOceanHTTPRequest `json:"http"`
}

func Main(ctx context.Context, params DigitalOceanParameters) (*DigitalOceanHTTPResponse, error) {
	fmt.Println(ctx)
	fmt.Println(params)
	return &DigitalOceanHTTPResponse{
		Body: fmt.Sprintf("Hello %s!", params.HTTP.QueryString),
	}, nil
}
