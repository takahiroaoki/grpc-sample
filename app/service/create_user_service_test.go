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

func TestCreateUserService_CreateUser_Success(t *testing.T) {
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

	mockRepository := mock.NewMockUserRepository(ctrl)
	mockRepository.EXPECT().CreateOneUser(db, u).Return(&entity.User{
		ID:    uint(1),
		Email: "user@example.com",
	}, nil)

	service := NewCreateUserService(mockRepository)
	actual, err := service.CreateUser(db, u)

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}

}

func TestCreateUserService_CreateUser_Error(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	u := entity.User{
		Email: "user@example.com",
	}
	var expected *entity.User

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockUserRepository(ctrl)
	mockRepository.EXPECT().CreateOneUser(db, u).Return(nil, util.NewError("err"))

	service := NewCreateUserService(mockRepository)
	actual, err := service.CreateUser(db, u)

	if assert.Error(t, err) {
		assert.Equal(t, "err", err.Error())
		assert.Equal(t, expected, actual)
	}

}
