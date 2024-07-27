package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/go-env/app/entity"
	"github.com/takahiroaoki/go-env/app/testutil"
	"github.com/takahiroaoki/go-env/app/testutil/mock"
	"github.com/takahiroaoki/go-env/app/util"
)

func TestSampleService_GetUserById_Success(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	userId := "1"
	expected := &entity.User{
		Email: "user@example.com",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockSampleRepository(ctrl)
	mockRepository.EXPECT().SelectOneUserByUserId(db, userId).Return(&entity.User{
		Email: "user@example.com",
	}, nil)

	sampleService := NewSampleService(mockRepository)
	actual, err := sampleService.GetUserByUserId(db, userId)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestSampleService_GetUserById_Error(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	userId := "1"
	var expected *entity.User

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockSampleRepository(ctrl)
	mockRepository.EXPECT().SelectOneUserByUserId(db, userId).Return(nil, util.NewError("err"))

	sampleService := NewSampleService(mockRepository)
	actual, err := sampleService.GetUserByUserId(db, userId)

	if assert.Error(t, err) {
		assert.Equal(t, "err", err.Error())
		assert.Equal(t, expected, actual)
	}
}

func TestSampleService_CreateUser_Success(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	u := entity.User{
		Email: "user@example.com",
	}
	expected := &entity.User{
		ID:    uint(1),
		Email: "user@example.com",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockSampleRepository(ctrl)
	mockRepository.EXPECT().CreateOneUser(db, u).Return(&entity.User{
		ID:    uint(1),
		Email: "user@example.com",
	}, nil)

	sampleService := NewSampleService(mockRepository)
	actual, err := sampleService.CreateUser(db, u)

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}

}

func TestSampleService_CreateUser_Error(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	u := entity.User{
		Email: "user@example.com",
	}
	var expected *entity.User

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockSampleRepository(ctrl)
	mockRepository.EXPECT().CreateOneUser(db, u).Return(nil, util.NewError("err"))

	sampleService := NewSampleService(mockRepository)
	actual, err := sampleService.CreateUser(db, u)

	if assert.Error(t, err) {
		assert.Equal(t, "err", err.Error())
		assert.Equal(t, expected, actual)
	}

}
