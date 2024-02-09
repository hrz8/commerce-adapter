package main

import (
	"context"
	"encoding/json"
	"fmt"
)

type DigitalOceanHTTPRequest struct {
	Headers         map[string]string `json:"headers"`
	Path            string            `json:"path"`
	Method          string            `json:"method"`
	Body            string            `json:"body"`
	QueryString     string            `json:"queryString"`
	IsBase64Encoded bool              `json:"isBase64Encoded"`
}

type DigitalOceanHTTPResponse struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

type DigitalOceanParameters struct {
	Headers map[string]string       `json:"__ow_headers"`
	Path    string                  `json:"__ow_path"`
	Method  string                  `json:"__ow_method"`
	Body    string                  `json:"__ow_body"`
	Query   string                  `json:"__ow_query"`
	HTTP    DigitalOceanHTTPRequest `json:"http"`
}

func Main(ctx context.Context, event DigitalOceanParameters) (*DigitalOceanHTTPResponse, error) {
	fmt.Println(fmt.Sprintf("params: %+v\n", event))
	jsonString, err := json.Marshal(event)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("JSON String:", string(jsonString))

	host := ctx.Value("api_host").(string)
	name := ctx.Value("function_name").(string)
	fmt.Println("ctx:", host, name)

	return &DigitalOceanHTTPResponse{
		Body: fmt.Sprintf("Hello %s!", "stranger"),
	}, nil
}
