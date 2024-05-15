package handlers

import (
	"github.com/ruziba3vich/events/repository"
)

type handler struct {
	services repository.UserRepository
}

type HandlerConfig struct {
	Services repository.UserRepository
}

func New(c *HandlerConfig) *handler {

	return &handler{
		services: c.Services,
	}
}
