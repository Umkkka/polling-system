package handler

import (
	"polling-system/internal/service"
)

type PollService interface {
	Create(poll *service.PollInfo) (string, error)
	Get(uuid string) (*service.PollInfo, error)
	SaveVote(uuid, answer string) error
}

func New(us PollService) *Handler {
	return &Handler{
		poll: us,
	}
}

type Handler struct {
	poll PollService
}
