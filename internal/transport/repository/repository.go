package repository

import "sync"

func New() *Repo {
	return &Repo{
		data: make(map[string]PollDTO),
		mu:   sync.Mutex{},
	}
}

type Repo struct {
	data map[string]PollDTO
	mu   sync.Mutex
}
