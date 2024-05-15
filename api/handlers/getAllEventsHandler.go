package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ruziba3vich/events/models"
)

func (s *handler) GetAllEventsHandler(c *gin.Context) {
	parsedUUID, err := uuid.Parse(c.Param("eventId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err = s.services.DeleteEvent(
		models.DeleteEventRequest{
			EventId: parsedUUID,
		})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, nil)
}
