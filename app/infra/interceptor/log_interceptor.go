package interceptor

import (
	"context"
	"fmt"
	"time"

	"github.com/takahiroaoki/grpc-sample/app/util"
	"google.golang.org/grpc"
)

func Log() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res any, err error) {
		reqAt := time.Now()
		util.InfoLogWithContext(ctx, fmt.Sprintf("Request: %v", info.FullMethod))

		defer func() {
			latency := time.Since(reqAt)

			if err != nil {
				util.ErrorLogWithContext(ctx, err.Error())
			}

			util.PerfLogWithContext(ctx, fmt.Sprintf("Response: %v, Latency: %v", info.FullMethod, latency))
		}()

		res, err = handler(ctx, req)
		return
	}
}
