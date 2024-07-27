package interceptor

import (
	"context"

	"github.com/takahiroaoki/go-env/app/interceptor/validator"
	"github.com/takahiroaoki/go-env/app/pb"
	"google.golang.org/grpc"
)

func ValidateReq() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		var err error
		switch req := req.(type) {
		case *pb.GetUserInfoRequest:
			err = validator.ValidateGetUserInfo(ctx, req)
		}

		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}
