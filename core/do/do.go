package docore

type DigitalOceanHTTPHeaderRequest struct {
	Accept          string `json:"accept"`
	AcceptEncoding  string `json:"accept-encoding"`
	ContentType     string `json:"content-type"`
	Host            string `json:"host"`
	UserAgent       string `json:"user-agent"`
	XForwardedFor   string `json:"x-forwarded-for"`
	XForwardedProto string `json:"x-forwarded-proto"`
	XRequestID      string `json:"x-request-id"`
}

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
