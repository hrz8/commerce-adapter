package main

import (
	"context"
	"encoding/json"
	"fmt"
)

type DigitalOceanHTTPRequest struct {
	Body            string            `json:"body"`
	Headers         map[string]string `json:"headers"`
	Method          string            `json:"method"`
	Path            string            `json:"path"`
	IsBase64Encoded bool              `json:"isBase64Encoded"`
	QueryString     string            `json:"queryString"`
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

	return &DigitalOceanHTTPResponse{
		Body: fmt.Sprintf("Hello %s!", "stranger"),
	}, nil
}
