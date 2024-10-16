package interceptor

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/infra/pb"
	"github.com/takahiroaoki/grpc-sample/app/infra/server/interceptor/validator"
	"google.golang.org/grpc"
)

func Validate() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		v := validator.NewValidator()
		var err error
		switch request := req.(type) {
		case *pb.CreateUserRequest:
			err = v.ValidateCreateUserRequest(ctx, request)
		case *pb.GetUserInfoRequest:
			err = v.ValidateGetUserInfoRequest(ctx, request)
		default:
		}
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}
