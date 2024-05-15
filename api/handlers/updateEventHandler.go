package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ruziba3vich/events/models"
)

func (s *handler) UpdateEventHandler(c *gin.Context) {
	eventId := c.Param("eventId")

	parsedUUID, err := uuid.Parse(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var req models.UpdateEventRequest
	req.Id = parsedUUID

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	event, err := s.services.UpdateEvent(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, event)
}
