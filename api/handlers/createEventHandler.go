package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/events/api/middleware"
	"github.com/ruziba3vich/events/models"
)

func (s *handler) CreateEventHandler(c *gin.Context) {
	userId, err := middleware.GetUserIdFromParam(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	var req models.CreateEventRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	req.UserId = userId

	event, err := s.services.CreateEvent(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, event)
}
