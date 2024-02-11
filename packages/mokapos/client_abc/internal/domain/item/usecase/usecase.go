package usecase

import (
	"aiconec/commerce-adapter/internal/port"
	"context"
)

type ItemUsecase struct {
	cfg port.ServiceConfig
}

func New(cfg port.ServiceConfig) *ItemUsecase {
	return &ItemUsecase{
		cfg,
	}
}

func (u *ItemUsecase) GetItems(ctx context.Context) (string, error) {
	return "Hello, World!", nil
}
