package cases

import (
	"fmt"
	"testing"

	"github.com/takahiroaoki/grpc-sample/tests/client"
	pb "github.com/takahiroaoki/protobuf/gen_go/proto/grpc_sample/v1"
)

func setup(t *testing.T) pb.SampleServiceClient {
	client, err := client.NewGRPCClient()
	if err != nil {
		fmt.Printf("failed to get client: %v", err)
		t.FailNow()
	}
	return client
}
