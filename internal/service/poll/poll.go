package poll

import (
	"context"
	"fmt"

	"polling-system/internal/service"
)

type Repository interface {
	Create(ctx context.Context, info *service.PollInfo) (uuid string, err error)
	Get(ctx context.Context, uuid string) (*service.PollInfo, error)
}

func New(repo Repository) *Poll {
	return &Poll{
		repo: repo,
	}
}

type Poll struct {
	repo Repository
}

func (p *Poll) Create(ctx context.Context, info *service.PollInfo) (string, error) {
	uuid, err := p.repo.Create(ctx, info)
	if err != nil {
		return "", fmt.Errorf("filed to create poll: %w", err)
	}

	return uuid, nil
}

func (p *Poll) Get(ctx context.Context, uuid string) (*service.PollInfo, error) {
	pollInfo, err := p.repo.Get(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to get poll: %w", err)
	}

	return pollInfo, nil
}
