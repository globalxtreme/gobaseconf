package client

import (
	"context"
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/grpc/pkg/bug"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"os"
	"time"
)

var (
	BugRPCClient bug.BugServiceClient
	BugRPCActive bool
)

func InitBugRPC() func() {
	addr := os.Getenv("GRPC_BUG_HOST")
	if !config.DevMode && addr != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		keepaliveParam := keepalive.ClientParameters{
			Time:                10 * time.Second,
			Timeout:             20 * time.Second,
			PermitWithoutStream: true,
		}

		conn, err := grpc.DialContext(ctx, addr,
			grpc.WithInsecure(),
			grpc.WithBlock(),
			grpc.WithKeepaliveParams(keepaliveParam),
		)
		if err != nil {
			log.Panicf("Did not connect to %s: %v", addr, err)
		}

		BugRPCClient = bug.NewBugServiceClient(conn)
		BugRPCActive = true

		cleanup := func() {
			cancel()
			conn.Close()
		}

		return cleanup
	}

	return func() {}
}

func BugLog(req *bug.LogRequest) (*bug.BGResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return BugRPCClient.Log(ctx, req)
}
