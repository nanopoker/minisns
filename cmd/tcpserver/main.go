package main

import (
	"flag"
	"fmt"
	"github.com/nanopoker/minisns/agent/cache"
	"github.com/nanopoker/minisns/agent/db"
	"github.com/nanopoker/minisns/apps/tcpserver"
	"github.com/nanopoker/minisns/config"
	"github.com/nanopoker/minisns/libs/logger"
	pb "github.com/nanopoker/minisns/proto"
	"google.golang.org/grpc"
	"net"
)

var (
	host string
	port string
)

func main() {
	flag.StringVar(&host, "host", config.TCP_HOST, "http server port")
	flag.StringVar(&port, "port", config.TCP_PORT, "http server port")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		logger.Error("failed to listen: %v", err)
	}

	err = db.Init()
	if err != nil {
		panic(err)
	}
	err = cache.Init()
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &tcpserver.Server{})
	logger.Info("starting to listen", fmt.Sprintf("%s:%s", host, port))
	grpcServer.Serve(lis)
}
