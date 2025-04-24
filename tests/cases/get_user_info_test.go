package cases

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/tests/pb"
)

func Test_CreateAndGetUser(t *testing.T) {
	client := setup(t)

	// call CreateUser
	createReq := &pb.CreateUserRequest{
		Email: "test@example.com",
	}
	gotCreateResp, err := client.CreateUser(context.Background(), createReq)
	if err != nil {
		fmt.Printf("failed to CreateUser: %v", err)
		t.FailNow()
	}

	// call GetUserInfo
	getReq := &pb.GetUserInfoRequest{
		Id: gotCreateResp.GetId(),
	}
	expectedGetResp := &pb.GetUserInfoResponse{
		Id:    gotCreateResp.GetId(),
		Email: "test@example.com",
	}
	gotGetResp, err := client.GetUserInfo(context.Background(), getReq)
	if err != nil {
		fmt.Printf("failed to GetUserInfo: %v", err)
	}
	assertGetUserInfoResponse(t, expectedGetResp, gotGetResp)
}

func Test_GetNonExistingUser(t *testing.T) {
	client := setup(t)
	getReq := &pb.GetUserInfoRequest{
		Id: "0",
	}
	gotGetResp, err := client.GetUserInfo(context.Background(), getReq)
	assert.Nil(t, gotGetResp)
	assert.EqualError(t, err, "rpc error: code = NotFound desc = data not found")
}
