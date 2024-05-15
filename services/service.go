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

func (s *Service) LogIn(req models.RegisterRequest) (string, error) {
	return s.userRepo.LogIn(req)
}

func (s *Service) CreateEvent(req models.CreateEventRequest) (*postgres.EventDTO, error) {
	return s.userRepo.CreateEvent(req)
}

func (s *Service) UpdateEvent(req models.UpdateEventRequest) (*postgres.EventDTO, error) {
	return s.userRepo.UpdateEvent(req)
}

func (s *Service) GetAllEvents(req models.GetAllEventsRequest) ([]postgres.EventDTO, error) {
	return s.userRepo.GetAllEvents(req)
}

func (s *Service) DeleteEvent(req models.DeleteEventRequest) error {
	return s.userRepo.DeleteEvent(req)
}
