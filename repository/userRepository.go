package repository

import "github.com/ruziba3vich/events/models"

type UserRepository interface {
	Register(models.RegisterRequest) (string, error)
	LogIn() error
	CreateEvent() error
	UpdateEvent() error
	DeleteEvent() error
	GetAllEvents() error
}

func NewUserRepository() {
	
}
