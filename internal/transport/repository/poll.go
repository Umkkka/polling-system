package repository

import (
	"context"

	"polling-system/internal/service"

	"github.com/google/uuid"
)

type PollDTO struct {
	Name string
}

func (r *Repo) Create(ctx context.Context, u *service.PollInfo) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	uuid := uuid.NewString()

	r.data[uuid] = PollDTO{
		Name: u.Name,
	}

	return uuid, nil
}
