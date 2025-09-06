package cases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	pb "github.com/takahiroaoki/protobuf/gen_go/proto/grpc_sample/v1"
)

func assertGetUserInfoResponse(t *testing.T, expected, got *pb.GetUserInfoResponse) {
	assert.Equal(t, expected.GetId(), got.GetId())
	assert.Equal(t, expected.GetEmail(), got.GetEmail())
}
