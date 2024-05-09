package poll

import (
	"context"
	"fmt"

	"polling-system/internal/service"
)

type Repository interface {
	Create(ctx context.Context, info *service.PollInfo) (uuid string, err error)
}

func New(repo Repository) *poll {
	return &poll{
		repo: repo,
	}
}

type poll struct {
	repo Repository
}

func (p *poll) Create(ctx context.Context, info *service.PollInfo) (string, error) {
	uuid, err := p.repo.Create(ctx, info)
	if err != nil {
		return "", fmt.Errorf("filed to create poll: %w", err)
	}

	return uuid, nil
}
