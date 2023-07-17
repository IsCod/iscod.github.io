package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginReq struct {
	Username string `form:"username"`
	Password string `form:"password" binding:"required"`
	Phone    string `form:"phone" binding:"required,e164"`
}

func Login(c *gin.Context) {
	req := &loginReq{}
	err := c.Bind(req) //模型绑定 c.BindJSON(req), c.BindQuery(req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"username": req.Username})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "register"})
}
