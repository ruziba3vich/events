package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id uuid.UUID `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	// Files []
	EventTime time.Time
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
