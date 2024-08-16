package handler

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/go-env/app/entity"
	"github.com/takahiroaoki/go-env/app/pb"
	"github.com/takahiroaoki/go-env/app/testutil"
	"github.com/takahiroaoki/go-env/app/testutil/mock"
	"github.com/takahiroaoki/go-env/app/util"
)

func TestCreateUserHandler_createUser_Success(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	ctx := context.Background()
	u := entity.User{
		Email: "user@example.com",
	}
	expected := &pb.CreateUserResponse{
		Id: "1",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockCreateUserService(ctrl)
	mockService.EXPECT().CreateUser(gomock.Any(), u).Return(&entity.User{
		ID:    uint(1),
		Email: "user@example.com",
	}, nil)

	handler := NewCreateUserHandler(db, mockService)
	actual, err := handler.createUser(ctx, &pb.CreateUserRequest{
		Email: u.Email,
	})
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestCreateUserHandler_createUser_Error(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	ctx := context.Background()
	u := entity.User{
		Email: "user@example.com",
	}
	var expected *pb.CreateUserResponse

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockCreateUserService(ctrl)
	mockService.EXPECT().CreateUser(gomock.Any(), u).Return(nil, util.NewError("err"))

	handler := NewCreateUserHandler(db, mockService)
	actual, err := handler.createUser(ctx, &pb.CreateUserRequest{
		Email: u.Email,
	})
	if assert.Error(t, err) {
		assert.Equal(t, "err", err.Error())
		assert.Equal(t, expected, actual)
	}
}
