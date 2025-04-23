package cases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/tests/pb"
)

func assertGetUserInfoResponse(t *testing.T, expected, got *pb.GetUserInfoResponse) {
	assert.Equal(t, expected.GetId(), got.GetId())
	assert.Equal(t, expected.GetEmail(), got.GetEmail())
}
