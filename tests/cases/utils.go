package cases

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/tests/client"
	"github.com/takahiroaoki/grpc-sample/tests/pb"
)

func setup(t *testing.T) pb.SampleServiceClient {
	client, err := client.NewGRPCClient()
	if err != nil {
		fmt.Printf("failed to get client: %v", err)
		t.FailNow()
	}
	return client
}

func assertGetUserInfoResponse(t *testing.T, expected, got *pb.GetUserInfoResponse) {
	assert.Equal(t, expected.GetId(), got.GetId())
	assert.Equal(t, expected.GetEmail(), got.GetEmail())
}
