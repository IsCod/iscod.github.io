package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

var (
	// 初始化为 unknown，如果编译时没有传入这些值，则为 unknown
	CommitID = "unknown"
)

var v = flag.Bool("v", false, "show version")

func getClientIp() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}

		}
	}

	return "", errors.New("Can not find the client ip address!!")

}

func main() {
	flag.Parse()
	if *v {
		_, _ = fmt.Println(CommitID)
		// testRandom()
		os.Exit(0)
	}
	c, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error creating %s", err)
		return
	}
	g := gin.New()
	g.Handle("GET", "/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
	g.Handle("GET", "/version", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("version: %s", CommitID)})
	})

	g.Handle("GET", "/ip", func(context *gin.Context) {
		var ip, _ = getClientIp()
		context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("IP: %s", ip)})
	})

	g.Use(gin.Logger(), gin.Recovery()) //添加logger和Recovery中间件

	if err = g.RunListener(c); err != nil {
		fmt.Printf("Error creating %s", err)
	}

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}
