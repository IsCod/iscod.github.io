package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iscod/example/handler"
	v22 "github.com/iscod/example/handler/v2"
	"github.com/iscod/example/middleware"
)

func initApi(g *gin.Engine) {
	api := g.Group("/api")
	v1 := api.Group("/v1")
	v1.Use(middleware.CheckToken, middleware.CheckAuth)
	v1.POST("/login", handler.Login)
	v1.POST("/register", handler.Register)

	v2 := api.Group("/v2")
	v2.POST("/login", v22.Login)
}

func InitRouter(g *gin.Engine) {
	initApi(g)
}
