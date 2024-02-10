package proxy

type ctxKey struct{}

const (
	defaultStatusCode    = -1
	contentTypeHeaderKey = "Content-Type"
	CustomHostVariable   = "GO_API_HOST"
	APIGwContextHeader   = "X-GoLambdaProxy-ApiGw-Context"
	APIGwStageVarsHeader = "X-GoLambdaProxy-ApiGw-StageVars"
)
