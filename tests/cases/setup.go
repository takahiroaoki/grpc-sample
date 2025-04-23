package cases

import (
	"fmt"
	"testing"

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
