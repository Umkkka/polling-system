package poll

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"

	"polling-system/internal/service"
)

type Repository interface {
	Create(poll *service.PollInfo) (uuid string, err error)
	Get(uuid string) (*service.PollInfo, error)
	SaveVote(uuid, answer string) (map[string]int, error)
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
	voteCount, err := p.repo.SaveVote(uuid, answer)
	if err != nil {
		return fmt.Errorf("failed to save vote: %w", err)
	}

	err = p.sendMessage(uuid, voteCount)
	if err != nil {
		return fmt.Errorf("failed to notify vote results: %w", err)
	}

	return nil
}

func (p *Poll) sendMessage(uuid string, voteCount map[string]int) error {
	poll, err := p.repo.Get(uuid)
	if err != nil {
		return fmt.Errorf("failed to get poll: %w", err)
	}

	message := "Poll Results:\n"
	message += "Poll ID: " + uuid + "\n"
	for _, option := range poll.Options {
		count := voteCount[option]
		message += "Option: " + option + ", Votes: " + strconv.Itoa(count) + "\n"
	}

	// message for websocket
	logrus.Info(message)

	err = p.sendPollResults(uuid, message)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

func (p *Poll) sendPollResults(uuid, message string) error {
	// TODO

	// Get websocket connection

	// Send message to websocket

	return nil
}
