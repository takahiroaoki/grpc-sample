package interceptor

import (
	"context"

	"github.com/takahiroaoki/go-env/app/interceptor/validator"
	"github.com/takahiroaoki/go-env/app/pb"
	"google.golang.org/grpc"
)

func ValidateReq() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		var err error
		switch req := req.(type) {
		case *pb.GetUserInfoRequest:
			err = validator.ValidateGetUserInfoRequest(ctx, req)
		case *pb.CreateUserRequest:
			err = validator.ValidateCreateUserRequest(ctx, req)
		}

		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}
