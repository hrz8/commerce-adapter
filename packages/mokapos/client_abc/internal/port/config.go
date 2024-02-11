package port

type ServiceConfig interface {
	GetBaseURL() string
	GetFunctionName() string
}
