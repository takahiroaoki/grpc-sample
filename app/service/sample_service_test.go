package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/go-env/entity"
	"github.com/takahiroaoki/go-env/mock"
	"github.com/takahiroaoki/go-env/util"
)

func TestSampleService_GetUserById_Success(t *testing.T) {
	t.Parallel()

	userId := "1"
	expected := &entity.User{
		Email: "user@example.com",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockSampleRepository(ctrl)
	mockRepository.EXPECT().SelectOneUserByUserId(userId).Return(&entity.User{
		Email: "user@example.com",
	}, nil)

	sampleService := NewSampleService(mockRepository)
	actual, err := sampleService.GetUserByUserId(userId)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestSampleService_GetUserById_Error(t *testing.T) {
	t.Parallel()

	userId := "1"
	var expected *entity.User

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockSampleRepository(ctrl)
	mockRepository.EXPECT().SelectOneUserByUserId(userId).Return(nil, util.NewError("err"))

	sampleService := NewSampleService(mockRepository)
	actual, err := sampleService.GetUserByUserId(userId)

	if assert.Error(t, err) {
		assert.Equal(t, "err", err.Error())
		assert.Equal(t, expected, actual)
	}
}
