package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iscod/example/router"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

// test curl
// curl --header 'token: adfsdf' --form 'username="iscod"' --form 'password="adfasdf"' --form 'phone="+8617091900050"' http://127.0.0.1:8080/api/v1/login

func main() {
	c, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("Error creating %s", err)
		return
	}
	g := gin.New()
	g.Handle("GET", "/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
	g.POST("/post", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	g.Use(gin.Logger(), gin.Recovery()) //添加logger和Recovery中间件

	router.InitRouter(g)

	if err = g.RunListener(c); err != nil {
		fmt.Printf("Error creating %s", err)
	}

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := c.Shutdown(ctx); err != nil {
	//	log.Fatal("Server Shutdown:", err)
	//}
	//log.Println("Server exiting")
}
