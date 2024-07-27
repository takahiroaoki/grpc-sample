package validator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/go-env/app/pb"
)

func TestValidateGetUserInfoRequest_Success(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	req := &pb.GetUserInfoRequest{
		Id: "12345",
	}

	var expected error
	actual := ValidateGetUserInfoRequest(ctx, req)

	assert.Equal(t, expected, actual)
}

func TestValidateGetUserInfoRequest_Error_Idが存在しない(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	req := &pb.GetUserInfoRequest{}

	expected := "id: cannot be blank."
	actual := ValidateGetUserInfoRequest(ctx, req)

	if assert.Error(t, actual) {
		assert.Equal(t, expected, actual.Error())
	}
}

func TestValidateGetUserInfoRequest_Error_Idが空文字(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	req := &pb.GetUserInfoRequest{
		Id: "",
	}

	expected := "id: cannot be blank."
	actual := ValidateGetUserInfoRequest(ctx, req)

	if assert.Error(t, actual) {
		assert.Equal(t, expected, actual.Error())
	}
}

func TestValidateGetUserInfoRequest_Error_Idが数字以外を含む(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	req := &pb.GetUserInfoRequest{
		Id: "abc",
	}

	expected := "id: must contain digits only."
	actual := ValidateGetUserInfoRequest(ctx, req)

	if assert.Error(t, actual) {
		assert.Equal(t, expected, actual.Error())
	}
}
