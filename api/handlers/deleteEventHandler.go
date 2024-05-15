package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ruziba3vich/events/models"
)

func (s *handler) DeleteEventHandler(c *gin.Context) {
	eventId := c.Param("eventId")

	parsedUUID, err := uuid.Parse(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = s.services.DeleteEvent(models.DeleteEventRequest{
		EventId: parsedUUID,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, nil)
}
