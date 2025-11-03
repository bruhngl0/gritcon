package repository

import "github.com/bruhngl0/gritcon/internal/server"

type Repositories struct{}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}
