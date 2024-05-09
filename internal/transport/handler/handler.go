package handler

import (
	"context"

	"polling-system/internal/service"
)

type PollService interface {
	Create(ctx context.Context, info *service.PollInfo) (string, error)
}

func New(us PollService) *Handler {
	return &Handler{
		poll: us,
	}
}

type Handler struct {
	poll PollService
}
