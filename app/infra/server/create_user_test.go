package server

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/handler"
	pb "github.com/takahiroaoki/protobuf/gen_go/proto/grpc_sample/v1"
	"github.com/takahiroaoki/grpc-sample/app/testutil/mockhandler"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_sampleServiceServer_CreateUser(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHandler := mockhandler.NewMockCreateUserHandler(ctrl)

	type args struct {
		ctx context.Context
		req *pb.CreateUserRequest
	}

	tests := []struct {
		name        string
		args        args
		mockFunc    func(ctx context.Context, mockHandler *mockhandler.MockCreateUserHandler)
		expected    *pb.CreateUserResponse
		assertion   assert.ErrorAssertionFunc
		expectedErr error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				req: &pb.CreateUserRequest{
					Email: "test@test.com",
				},
			},
			mockFunc: func(ctx context.Context, mockHandler *mockhandler.MockCreateUserHandler) {
				mockHandler.EXPECT().Invoke(ctx, handler.NewCreateUserRequest("test@test.com")).Return(handler.NewCreateUserResponse("id"), nil)
			},
			expected: &pb.CreateUserResponse{
				Id: "id",
			},
			assertion: assert.NoError,
		},
		{
			name: "Error",
			args: args{
				ctx: context.Background(),
				req: &pb.CreateUserRequest{
					Email: "test@test.com",
				},
			},
			mockFunc: func(ctx context.Context, mockHandler *mockhandler.MockCreateUserHandler) {
				mockHandler.EXPECT().Invoke(ctx, handler.NewCreateUserRequest("test@test.com")).Return(nil, domerr.NewDomErrFromMsg("error", domerr.CAUSE_INTERNAL, domerr.LOG_LEVEL_ERROR))
			},
			expected:    nil,
			assertion:   assert.Error,
			expectedErr: status.Error(codes.Internal, "internal error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grpcServiceServer := &sampleServiceServer{
				createUserHandler: mockHandler,
			}
			if tt.mockFunc != nil {
				tt.mockFunc(tt.args.ctx, mockHandler)
			}
			res, err := grpcServiceServer.CreateUser(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.expected, res)
			tt.assertion(t, err)
			if err != nil {
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}
