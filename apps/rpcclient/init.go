package rpcclient

import (
	"context"
	"fmt"
	"github.com/nanopoker/minisns/config"
	"github.com/nanopoker/minisns/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

var conn *grpc.ClientConn
var client proto.UserServiceClient
var ctx context.Context

func init() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", config.TCP_HOST, config.TCP_PORT), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client = proto.NewUserServiceClient(conn)
	ctx, _ = context.WithTimeout(context.Background(), 7200*time.Second)
}
