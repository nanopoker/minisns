package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/nanopoker/minisns/apps/routers"
	"github.com/nanopoker/minisns/config"
	"github.com/nanopoker/minisns/libs/logger"
	"os"
)

var (
	host string
	port string
)

func main() {
	flag.StringVar(&host, "host", config.HTTP_HOST, "http server port")
	flag.StringVar(&port, "port", config.HTTP_PORT, "http server port")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	server.Use(gin.Recovery())

	gin.DefaultWriter = logger.LogFile

	router.InitializeRouter(server)
	logger.Info("runServer start working...")
	if err := server.Run(host + ":" + port); err != nil {
		logger.Error("init server error: %v---", err)
		os.Exit(-1)
	}
}
