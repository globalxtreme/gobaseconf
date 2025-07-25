package client

import (
	"context"
	"github.com/globalxtreme/gobaseconf/config"
	log2 "github.com/globalxtreme/gobaseconf/grpc/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	// LogRPCClient --> Log service gRPC client
	LogRPCClient log2.LogServiceClient

	// LogRPCTimeout --> Log service gRPC timeout while send log
	LogRPCTimeout time.Duration

	// LogRPCActive --> Log service gRPC status active or inactive
	LogRPCActive bool
)

func InitLogRPC(force ...bool) func() {
	isForce := false
	if len(force) > 0 {
		isForce = force[0]
	}

	addr := os.Getenv("GRPC_LOG_HOST")
	if (isForce || !config.DevMode) && addr != "" {
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

		LogRPCClient = log2.NewLogServiceClient(conn)
		LogRPCActive = true

		LogRPCTimeout = 5 * time.Second
		if bugTimeoutENV := os.Getenv("GRPC_LOG_TIMEOUT"); bugTimeoutENV != "" {
			bugTimeoutENVInt, _ := strconv.Atoi(bugTimeoutENV)

			LogRPCTimeout = time.Duration(bugTimeoutENVInt) * time.Second
		}

		cleanup := func() {
			cancel()
			conn.Close()
		}

		return cleanup
	}

	return func() {}
}
