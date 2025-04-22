package client

import (
	"context"
	"errors"
	"time"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/takahiroaoki/grpc-sample/tests/config"
	"github.com/takahiroaoki/grpc-sample/tests/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCClient() (pb.SampleServiceClient, error) {
	clientConfig := config.NewClientConfig()
	conn, err := connect(10, clientConfig)
	if err != nil {
		return nil, err
	}
	return pb.NewSampleServiceClient(conn), nil
}

func connect(retry int, config config.ClientConfig) (*grpc.ClientConn, error) {
	if retry <= 0 {
		return nil, errors.New("retry limit")
	}
	cc, err := grpc.NewClient(
		config.TargetAddress(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(
			middleware.ChainUnaryClient(
				UnaryTimeoutInterceptor(config.RequestTimeout()),
			),
		),
		grpc.WithStreamInterceptor(
			middleware.ChainStreamClient(
				StreamTimeoutInterceptor(config.RequestTimeout()),
			),
		),
	)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.ConnectTimeout())
	defer cancel()
	if err := connectWithContext(ctx, cc); err != nil {
		time.Sleep(config.ConnectRetryWait())
		return connect(retry-1, config)
	}
	return cc, nil
}

func connectWithContext(ctx context.Context, cc *grpc.ClientConn) error {
	for {
		s := cc.GetState()
		if s == connectivity.Idle {
			cc.Connect()
		}
		if s == connectivity.Ready {
			return nil
		}
		if !cc.WaitForStateChange(ctx, s) {
			return ctx.Err()
		}
	}
}
