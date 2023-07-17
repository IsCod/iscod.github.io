package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func CheckToken(c *gin.Context) {
	if c.Request.Header.Get("token") == "" {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Invalid username"})
		c.AbortWithError(http.StatusServiceUnavailable, errors.New("Invalid usernames"))
	}
	c.Next()
}

func CheckAuth(c *gin.Context) {
	if c.Request.PostFormValue("username") == "" {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Invalid username required"})
		c.AbortWithError(http.StatusServiceUnavailable, errors.New("Invalid username required"))
	}
	c.Next()
}
