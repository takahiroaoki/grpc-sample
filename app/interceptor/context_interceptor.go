package interceptor

import (
	"context"

	"github.com/google/uuid"
	"github.com/takahiroaoki/go-env/app/constant"
	"google.golang.org/grpc"
)

func SetContext() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res any, err error) {
		ctx = context.WithValue(ctx, constant.REQUEST_ID, uuid.New())
		res, err = handler(ctx, req)
		return
	}
}
