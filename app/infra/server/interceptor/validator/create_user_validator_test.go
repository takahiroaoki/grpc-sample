package validator

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	pb "github.com/takahiroaoki/protobuf/gen_go/proto/grpc_sample/v1"
)

func Test_Validate_ValidateCreateUserRequest(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx context.Context
		req *pb.CreateUserRequest
	}
	tests := []struct {
		name           string
		args           args
		assertion      assert.ErrorAssertionFunc
		expectedErrMsg string
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				req: &pb.CreateUserRequest{
					Email: "user@example.com",
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "Success(Email right boundary safe)",
			args: args{
				ctx: context.Background(),
				req: &pb.CreateUserRequest{
					Email: strings.Repeat("a", 308) + "@example.com",
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "Error(Email right boundary over)",
			args: args{
				ctx: context.Background(),
				req: &pb.CreateUserRequest{
					Email: strings.Repeat("a", 309) + "@example.com",
				},
			},
			assertion:      assert.Error,
			expectedErrMsg: "email: the length must be between 1 and 320.",
		},
		{
			name: "Error(Email is nil)",
			args: args{
				ctx: context.Background(),
				req: &pb.CreateUserRequest{},
			},
			assertion:      assert.Error,
			expectedErrMsg: "email: cannot be blank.",
		},
		{
			name: "Error(Email is empty)",
			args: args{
				ctx: context.Background(),
				req: &pb.CreateUserRequest{
					Email: "",
				},
			},
			assertion:      assert.Error,
			expectedErrMsg: "email: cannot be blank.",
		},
		{
			name: "Error(Email is in an invalid format)",
			args: args{
				ctx: context.Background(),
				req: &pb.CreateUserRequest{
					Email: "invalid format",
				},
			},
			assertion:      assert.Error,
			expectedErrMsg: "email: must be in a valid format.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewValidator()
			err := v.ValidateCreateUserRequest(tt.args.ctx, tt.args.req)
			tt.assertion(t, err)
			if err != nil {
				assert.Equal(t, tt.expectedErrMsg, err.Error())
			}
		})
	}
}
