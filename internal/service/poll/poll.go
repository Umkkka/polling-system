package poll

import (
	"fmt"

	"polling-system/internal/service"
)

type Repository interface {
	Create(poll *service.PollInfo) (uuid string, err error)
	Get(uuid string) (*service.PollInfo, error)
	SaveVote(uuid, answer string) error
}

func New(repo Repository) *Poll {
	return &Poll{
		repo: repo,
	}
}

type Poll struct {
	repo Repository
}

func (p *Poll) Create(poll *service.PollInfo) (string, error) {
	uuid, err := p.repo.Create(poll)
	if err != nil {
		return "", fmt.Errorf("filed to create poll: %w", err)
	}

	return uuid, nil
}

func (p *Poll) Get(uuid string) (*service.PollInfo, error) {
	pollInfo, err := p.repo.Get(uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to get poll: %w", err)
	}

	return pollInfo, nil
}

func (p *Poll) SaveVote(uuid, answer string) error {
	err := p.repo.SaveVote(uuid, answer)
	if err != nil {
		return fmt.Errorf("failed to save vote: %w", err)
	}

	return nil
}
