package main

import (
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
	// HTTP DigitalOceanHTTPRequest `json:"http"`
	Name string `json:"name"`
}

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

func Main(in DigitalOceanParameters) (*Response, error) {
	if in.Name == "" {
		in.Name = "stranger"
	}

	fmt.Println(fmt.Sprintf("params: %+v\n", in))
	// fmt.Println(in.HTTP.Method)
	// fmt.Println(in.HTTP.Path)

	return &Response{
		Body: fmt.Sprintf("Hello %s!", in.Name),
	}, nil
}
