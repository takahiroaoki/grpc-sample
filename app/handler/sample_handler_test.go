package handler

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/go-env/entity"
	"github.com/takahiroaoki/go-env/mock"
	"github.com/takahiroaoki/go-env/pb"
	"github.com/takahiroaoki/go-env/util"
	"gorm.io/gorm"
)

func TestSampleHander_GetUserInfo_Success(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	userId := "1"
	expected := &pb.GetUserInfoResponse{
		Id:    "1",
		Email: "user@example.com",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockSampleService(ctrl)
	mockService.EXPECT().GetUserByUserId(userId).Return(&entity.User{
		Model: gorm.Model{
			ID: uint(1),
		},
		Email: "user@example.com",
	}, nil)

	sampleHandler := NewSampleHandler(mockService)
	actual, err := sampleHandler.GetUserInfo(ctx, &pb.GetUserInfoRequest{
		Id: userId,
	})
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestSampleHander_GetUserInfo_Error(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	userId := "1"
	var expected *pb.GetUserInfoResponse

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockSampleService(ctrl)
	mockService.EXPECT().GetUserByUserId(userId).Return(nil, util.NewError("err"))

	sampleHandler := NewSampleHandler(mockService)
	actual, err := sampleHandler.GetUserInfo(ctx, &pb.GetUserInfoRequest{
		Id: userId,
	})
	if assert.Error(t, err) {
		assert.Equal(t, "err", err.Error())
		assert.Equal(t, expected, actual)
	}
}
