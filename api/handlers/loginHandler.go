package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/events/models"
)

func (s *handler) LogInHandler(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	token, err := s.services.LogIn(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
