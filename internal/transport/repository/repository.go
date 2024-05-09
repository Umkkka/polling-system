package repository

import "sync"

func New() *Repo {
	return &Repo{
		data:    make(map[string]PollDTO),
		results: make(map[string]map[string]int),
		mu:      sync.Mutex{},
	}
}

type Repo struct {
	data    map[string]PollDTO
	results map[string]map[string]int
	mu      sync.Mutex
}
