package services

import (
	"database/sql"

	"github.com/ruziba3vich/events/models"
	"github.com/ruziba3vich/events/postgres"
	"github.com/ruziba3vich/events/repository"
)

type Service struct {
	userRepo repository.UserRepository
}

func NewService(db *sql.DB) *Service {
	return &Service{
		userRepo: postgres.NewUserRepo(db),
	}
}

func (s *Service) Register(req models.RegisterRequest) (string, error) {
	return s.userRepo.Register(req)
}
