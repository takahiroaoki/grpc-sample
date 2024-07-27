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

func TestSampleHander_GetUserInfo_Success(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	ctx := context.Background()
	userId := "1"
	expected := &pb.GetUserInfoResponse{
		Id:    "1",
		Email: "user@example.com",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockSampleService(ctrl)
	mockService.EXPECT().GetUserByUserId(db, userId).Return(&entity.User{
		ID:    uint(1),
		Email: "user@example.com",
	}, nil)

	sampleHandler := NewSampleHandler(db, mockService)
	actual, err := sampleHandler.GetUserInfo(ctx, &pb.GetUserInfoRequest{
		Id: userId,
	})
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestSampleHander_GetUserInfo_Error(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	ctx := context.Background()
	userId := "1"
	var expected *pb.GetUserInfoResponse

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockSampleService(ctrl)
	mockService.EXPECT().GetUserByUserId(db, userId).Return(nil, util.NewError("err"))

	sampleHandler := NewSampleHandler(db, mockService)
	actual, err := sampleHandler.GetUserInfo(ctx, &pb.GetUserInfoRequest{
		Id: userId,
	})
	if assert.Error(t, err) {
		assert.Equal(t, "err", err.Error())
		assert.Equal(t, expected, actual)
	}
}

func TestSampleHander_CreateUser_Success(t *testing.T) {
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

	mockService := mock.NewMockSampleService(ctrl)
	mockService.EXPECT().CreateUser(gomock.Any(), u).Return(&entity.User{
		ID:    uint(1),
		Email: "user@example.com",
	}, nil)

	sampleHandler := NewSampleHandler(db, mockService)
	actual, err := sampleHandler.CreateUser(ctx, &pb.CreateUserRequest{
		Email: u.Email,
	})
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestSampleHander_CreateUser_Error(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	ctx := context.Background()
	u := entity.User{
		Email: "user@example.com",
	}
	var expected *pb.CreateUserResponse

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockSampleService(ctrl)
	mockService.EXPECT().CreateUser(gomock.Any(), u).Return(nil, util.NewError("err"))

	sampleHandler := NewSampleHandler(db, mockService)
	actual, err := sampleHandler.CreateUser(ctx, &pb.CreateUserRequest{
		Email: u.Email,
	})
	if assert.Error(t, err) {
		assert.Equal(t, "err", err.Error())
		assert.Equal(t, expected, actual)
	}
}
