package port

import "context"

type ServiceUsecase interface {
	GetItems(ctx context.Context) (string, error)
}
