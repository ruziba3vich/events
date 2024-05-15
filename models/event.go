package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id          uuid.UUID `json:"id"`
	UserId      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EventTime   time.Time `json:"event_time"`
	Files       []string  `json:"files"`
}

/*
- event
{
- id
- title
- description
- files
- event_time (timestamp)
}
*/
