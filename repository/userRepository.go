package repository

import (
	"github.com/ruziba3vich/events/models"
	"github.com/ruziba3vich/events/postgres"
)

type UserRepository interface {
	Register(models.RegisterRequest) (string, error)
	LogIn(models.RegisterRequest) (string, error)
	CreateEvent(models.CreateEventRequest) (*models.Event, error)
	UpdateEvent(models.UpdateEventRequest) (*postgres.EventDTO, error)
	DeleteEvent(models.DeleteEventRequest) error
	GetAllEvents(models.GetAllEventsRequest) ([]postgres.EventDTO, error)
}
