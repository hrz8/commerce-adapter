package main

import (
	"context"
	"encoding/json"
	"fmt"
)

type DigitalOceanHTTPRequest struct {
	Body            string              `json:"body"`
	Headers         map[string][]string `json:"headers"`
	IsBase64Encoded bool                `json:"isBase64Encoded,omitempty"`
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
	Body   string `json:"__ow_body"`
	Method string `json:"__ow_method"`
	Query  string `json:"__ow_query"`
	Name   string `json:"name"`
}

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

func Main(ctx context.Context, event DigitalOceanParameters) (*Response, error) {
	jsonString, err := json.Marshal(event)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(fmt.Sprintf("params: %s", jsonString))

	if event.Name == "" {
		event.Name = "stranger"
	}

	// fmt.Println(in.HTTP.Method)
	// fmt.Println(in.HTTP.Path)

	return &Response{
		Body: fmt.Sprintf("Hello %s!", event.Name),
	}, nil
}
