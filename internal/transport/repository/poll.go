package repository

import (
	"errors"

	"github.com/google/uuid"

	"polling-system/internal/service"
)

type PollDTO struct {
	Title   string
	Options []string
}

func (r *Repo) Create(poll *service.PollInfo) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	uuid := uuid.NewString()

	r.data[uuid] = PollDTO{
		Title:   poll.Title,
		Options: poll.Options,
	}
	r.results[uuid] = make(map[string]int)

	for _, option := range poll.Options {
		r.results[uuid][option] = 0
	}

	return uuid, nil
}

func (r *Repo) Get(uuid string) (*service.PollInfo, error) {
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

func (r *Repo) SaveVote(uuid, answer string) (voteCount map[string]int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	poll, ok := r.data[uuid]
	if !ok {
		return nil, errors.New("poll not found")
	}

	if !isExist(poll.Options, answer) {
		return nil, errors.New("answer does not exist in poll")
	}

	r.results[uuid][answer]++
	return r.results[uuid], nil
}

func isExist(arr []string, target string) bool {
	for _, s := range arr {
		if s == target {
			return true
		}
	}

	return false
}
