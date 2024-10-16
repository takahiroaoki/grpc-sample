package validator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/infra/pb"
)

func Test_Validator_ValidatepGetUserInfoRequest(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		req *pb.GetUserInfoRequest
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
				req: &pb.GetUserInfoRequest{
					Id: "12345",
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "Error(Id is nil)",
			args: args{
				ctx: context.Background(),
				req: &pb.GetUserInfoRequest{},
			},
			assertion:      assert.Error,
			expectedErrMsg: "id: cannot be blank.",
		},
		{
			name: "Error(Id is empty)",
			args: args{
				ctx: context.Background(),
				req: &pb.GetUserInfoRequest{
					Id: "",
				},
			},
			assertion:      assert.Error,
			expectedErrMsg: "id: cannot be blank.",
		},
		{
			name: "Error(Id contains invalid characters)",
			args: args{
				ctx: context.Background(),
				req: &pb.GetUserInfoRequest{
					Id: "abc",
				},
			},
			assertion:      assert.Error,
			expectedErrMsg: "id: must contain digits only.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewValidator()
			err := v.ValidateGetUserInfoRequest(tt.args.ctx, tt.args.req)
			tt.assertion(t, err)
			if err != nil {
				assert.Equal(t, tt.expectedErrMsg, err.Error())
			}
		})
	}
}
