package client

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

func UnaryTimeoutInterceptor(timeout time.Duration) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		newCtx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return invoker(newCtx, method, req, reply, cc, opts...)
	}
}

func StreamTimeoutInterceptor(timeout time.Duration) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		newCtx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return streamer(newCtx, desc, cc, method, opts...)
	}
}
