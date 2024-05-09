package repository

import (
	"context"
	"errors"

	"polling-system/internal/service"

	"github.com/google/uuid"
)

type PollDTO struct {
	Title   string
	Options []string
}

func (r *Repo) Create(ctx context.Context, p *service.PollInfo) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	uuid := uuid.NewString()

	r.data[uuid] = PollDTO{
		Title:   p.Title,
		Options: p.Options,
	}

	return uuid, nil
}

func (r *Repo) Get(ctx context.Context, uuid string) (*service.PollInfo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	pollDTO, ok := r.data[uuid]
	if !ok {
		return nil, errors.New("poll not found")
	}

	pollInfo := &service.PollInfo{
		Title:   pollDTO.Title,
		Options: pollDTO.Options,
	}

	return pollInfo, nil
}
