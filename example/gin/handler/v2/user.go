package v2

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {

}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "login"})
}
