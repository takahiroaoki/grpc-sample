package validator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/go-env/pb"
)

func TestGetUserInfoValidator_Success(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	req := &pb.GetUserInfoRequest{
		Id: "12345",
	}

	var expected error
	actual := ValidateGetUserInfo(ctx, req)

	assert.Equal(t, expected, actual)
}

func TestGetUserInfoValidator_Error_Id項目が存在しない(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	req := &pb.GetUserInfoRequest{}

	expected := "id: cannot be blank."
	actual := ValidateGetUserInfo(ctx, req).Error()

	assert.Equal(t, expected, actual)
}

func TestGetUserInfoValidator_Error_Idが空文字(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	req := &pb.GetUserInfoRequest{
		Id: "",
	}

	expected := "id: cannot be blank."
	actual := ValidateGetUserInfo(ctx, req).Error()

	assert.Equal(t, expected, actual)
}

func TestGetUserInfoValidator_Error_Idが数字以外を含む(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	req := &pb.GetUserInfoRequest{
		Id: "abc",
	}

	expected := "id: must contain digits only."
	actual := ValidateGetUserInfo(ctx, req).Error()

	assert.Equal(t, expected, actual)
}
