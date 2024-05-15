package models

import (
	"time"

	"github.com/google/uuid"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateEventRequest struct {
	UserId      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EventTime   time.Time `json:"event_time"`
	Files       []string  `json:"files"`
}

type UpdateEventRequest struct {
	Id          uuid.UUID `json:"event_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EventTime   time.Time `json:"event_time"`
}

type DeleteEventRequest struct {
	EventId uuid.UUID `json:"event_id"`
}

type GetAllEventsRequest struct {
	UserId int `json:"user_id"`
}

/*
type Event struct {
	Id          uuid.UUID `json:"id"`
	UserId      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EventTime   time.Time `json:"event_time"`
	Files       []string  `json:"files"`
}
*/

/*
type Event struct {
	Id uuid.UUID `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	EventTime time.Time
}
*/
