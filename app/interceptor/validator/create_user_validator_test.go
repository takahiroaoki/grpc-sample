package validator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/go-env/app/pb"
)

func TestValidateCreateUserRequest_Success(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	req := &pb.CreateUserRequest{
		Email: "user@example.com",
	}

	var expected error
	actual := ValidateCreateUserRequest(ctx, req)

	assert.Equal(t, expected, actual)
}

func TestValidateCreateUserRequest_Error_Emailが存在しない(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	req := &pb.CreateUserRequest{}

	expected := "email: cannot be blank."
	actual := ValidateCreateUserRequest(ctx, req)

	if assert.Error(t, actual) {
		assert.Equal(t, expected, actual.Error())
	}
}

func TestValidateCreateUserRequest_Error_Email項目が空文字(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	req := &pb.CreateUserRequest{
		Email: "",
	}

	expected := "email: cannot be blank."
	actual := ValidateCreateUserRequest(ctx, req)

	if assert.Error(t, actual) {
		assert.Equal(t, expected, actual.Error())
	}
}

func TestValidateCreateUserRequest_Error_Email項目が桁数オーバー(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	req := &pb.CreateUserRequest{
		// 321
		Email: "123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789@example.com",
	}

	expected := "email: the length must be between 1 and 320."
	actual := ValidateCreateUserRequest(ctx, req)

	if assert.Error(t, actual) {
		assert.Equal(t, expected, actual.Error())
	}
}

func TestValidateCreateUserRequest_Error_Email項目がフォーマットエラー(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	req := &pb.CreateUserRequest{
		Email: "format error",
	}

	expected := "email: must be in a valid format."
	actual := ValidateCreateUserRequest(ctx, req)

	if assert.Error(t, actual) {
		assert.Equal(t, expected, actual.Error())
	}
}
