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

func TestGetUserInfoHandler_getUserInfo_Success(t *testing.T) {
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

	mockService := mock.NewMockGetUserInfoService(ctrl)
	mockService.EXPECT().GetUserByUserId(db, userId).Return(&entity.User{
		ID:    uint(1),
		Email: "user@example.com",
	}, nil)

	handler := NewGetUserInfoHandler(db, mockService)
	actual, err := handler.getUserInfo(ctx, &pb.GetUserInfoRequest{
		Id: userId,
	})
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestGetUserInfoHandler_getUserInfo_Error(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	ctx := context.Background()
	userId := "1"
	var expected *pb.GetUserInfoResponse

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockGetUserInfoService(ctrl)
	mockService.EXPECT().GetUserByUserId(db, userId).Return(nil, util.NewError("err"))

	handler := NewGetUserInfoHandler(db, mockService)
	actual, err := handler.getUserInfo(ctx, &pb.GetUserInfoRequest{
		Id: userId,
	})
	if assert.Error(t, err) {
		assert.Equal(t, "err", err.Error())
		assert.Equal(t, expected, actual)
	}
}

func TestGetUserInfoHandler_validate_Success(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockGetUserInfoService(ctrl)

	handler := &getUserInfoHandlerImpl{
		db:                 db,
		getUserInfoService: mockService,
	}
	err := handler.validate(ctx, &pb.GetUserInfoRequest{
		Id: "12345",
	})
	assert.NoError(t, err)
}

func TestGetUserInfoHandler_validate_Error_Id項目が存在しない(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockGetUserInfoService(ctrl)

	expected := "id: cannot be blank."
	handler := &getUserInfoHandlerImpl{
		db:                 db,
		getUserInfoService: mockService,
	}
	err := handler.validate(ctx, &pb.GetUserInfoRequest{})
	if assert.Error(t, err) {
		assert.Equal(t, expected, err.Error())
	}
}

func TestGetUserInfoHandler_validate_Error_Idが空文字(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockGetUserInfoService(ctrl)

	expected := "id: cannot be blank."
	handler := &getUserInfoHandlerImpl{
		db:                 db,
		getUserInfoService: mockService,
	}
	err := handler.validate(ctx, &pb.GetUserInfoRequest{
		Id: "",
	})
	if assert.Error(t, err) {
		assert.Equal(t, expected, err.Error())
	}
}

func TestGetUserInfoHandler_validate_Error_Idが数字以外を含む(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockGetUserInfoService(ctrl)

	expected := "id: must contain digits only."
	handler := &getUserInfoHandlerImpl{
		db:                 db,
		getUserInfoService: mockService,
	}
	err := handler.validate(ctx, &pb.GetUserInfoRequest{
		Id: "abc",
	})
	if assert.Error(t, err) {
		assert.Equal(t, expected, err.Error())
	}
}
