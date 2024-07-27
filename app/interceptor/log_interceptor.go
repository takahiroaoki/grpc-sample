package interceptor

import (
	"context"
	"fmt"
	"time"

	"github.com/takahiroaoki/go-env/util"
	"google.golang.org/grpc"
)

func Log() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res any, err error) {
		reqAt := time.Now()
		util.InfoLog(fmt.Sprintf("Request: %v", info.FullMethod))

		defer func() {
			resAt := time.Now()
			latency := resAt.Sub(reqAt)

			if err != nil {
				util.ErrorLog(err.Error())
			}

			util.PerfLog(fmt.Sprintf("Response: %v, Latency: %v", info.FullMethod, latency))
		}()

		res, err = handler(ctx, req)
		return
	}
}
