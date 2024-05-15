package middleware

import (
	"github.com/gin-gonic/gin"
	auth "github.com/ruziba3vich/authentication_tokens"
)

func GetUserIdFromParam(c *gin.Context) (int, error) {
	token := c.GetHeader("Authorization")[7:]
	id, err := auth.ExtractUserIDFromToken(token)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetUsernameFromParam(c *gin.Context) (string, error) {
	token := c.GetHeader("Authorization")[7:]
	username, err := auth.ExtractUsernameFromToken(token)
	if err != nil {
		return "", err
	}
	return username, nil
}
