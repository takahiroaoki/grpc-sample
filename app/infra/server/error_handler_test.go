package server

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_handleError(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx    context.Context
		domerr domerr.DomErr
	}

	tests := []struct {
		name     string
		args     args
		expected error
	}{
		{
			name: "not found",
			args: args{
				ctx:    context.Background(),
				domerr: domerr.NewDomErrFromMsg("err", domerr.CAUSE_NOT_FOUND, domerr.LOG_LEVEL_INFO),
			},
			expected: status.Error(codes.NotFound, "data not found"),
		},
		{
			name: "invalid argument",
			args: args{
				ctx:    context.Background(),
				domerr: domerr.NewDomErrFromMsg("err", domerr.CAUSE_INVALID_ARGUMENT, domerr.LOG_LEVEL_INFO),
			},
			expected: status.Error(codes.InvalidArgument, "invalid argument"),
		},
		{
			name: "internal",
			args: args{
				ctx:    context.Background(),
				domerr: domerr.NewDomErrFromMsg("err", domerr.CAUSE_INTERNAL, domerr.LOG_LEVEL_INFO),
			},
			expected: status.Error(codes.Internal, "internal error"),
		},
		{
			name: "undefined",
			args: args{
				ctx:    context.Background(),
				domerr: domerr.NewDomErrFromMsg("err", domerr.CAUSE_UNDEFINED, domerr.LOG_LEVEL_INFO),
			},
			expected: status.Error(codes.Unknown, "unknown error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, handleError(tt.args.ctx, tt.args.domerr))
		})
	}
}
