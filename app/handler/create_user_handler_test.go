package handler

import (
	"context"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/go-env/app/entity"
	"github.com/takahiroaoki/go-env/app/pb"
	"github.com/takahiroaoki/go-env/app/testutil"
	"github.com/takahiroaoki/go-env/app/testutil/mock"
	"github.com/takahiroaoki/go-env/app/util"
)

func TestCreateUserHandler_execute_Success(t *testing.T) {
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
	actual, err := handler.execute(ctx, &pb.CreateUserRequest{
		Email: u.Email,
	})
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestCreateUserHandler_execute_Error_service(t *testing.T) {
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
	actual, err := handler.execute(ctx, &pb.CreateUserRequest{
		Email: u.Email,
	})
	if assert.Error(t, err) {
		assert.Equal(t, "err", err.Error())
		assert.Equal(t, expected, actual)
	}
}

func TestCreateUserHandler_execute_Error_validation(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	ctx := context.Background()
	u := entity.User{
		Email: "invalid value",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockCreateUserService(ctrl)
	mockService.EXPECT().CreateUser(gomock.Any(), gomock.Any()).MaxTimes(0)

	handler := NewCreateUserHandler(db, mockService)
	actual, err := handler.execute(ctx, &pb.CreateUserRequest{
		Email: u.Email,
	})
	assert.Nil(t, actual)
	assert.Error(t, err)
}

func TestCreateUserHandler_validate_Success(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockCreateUserService(ctrl)

	handler := &createUserHandler{
		db:                db,
		createUserService: mockService,
	}
	err := handler.validate(ctx, &pb.CreateUserRequest{
		Email: "user@example.com",
	})
	assert.NoError(t, err)
}

func TestCreateUserHandler_validate_Success_Emailの桁数境界値(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockCreateUserService(ctrl)

	handler := &createUserHandler{
		db:                db,
		createUserService: mockService,
	}
	err := handler.validate(ctx, &pb.CreateUserRequest{
		// 320桁
		Email: strings.Repeat("a", 308) + "@example.com",
	})
	assert.NoError(t, err)
}

func TestCreateUserHandler_validate_Error_Email項目が存在しない(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockCreateUserService(ctrl)

	handler := &createUserHandler{
		db:                db,
		createUserService: mockService,
	}
	expected := "email: cannot be blank."
	err := handler.validate(ctx, &pb.CreateUserRequest{})
	if assert.Error(t, err) {
		assert.Equal(t, expected, err.Error())
	}
}

func TestCreateUserHandler_validate_Error_Emailが空文字(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockCreateUserService(ctrl)

	handler := &createUserHandler{
		db:                db,
		createUserService: mockService,
	}
	expected := "email: cannot be blank."
	err := handler.validate(ctx, &pb.CreateUserRequest{
		Email: "",
	})
	if assert.Error(t, err) {
		assert.Equal(t, expected, err.Error())
	}
}

func TestCreateUserHandler_validate_Error_Emailが桁数オーバー(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockCreateUserService(ctrl)

	handler := &createUserHandler{
		db:                db,
		createUserService: mockService,
	}
	expected := "email: the length must be between 1 and 320."
	err := handler.validate(ctx, &pb.CreateUserRequest{
		// 321桁
		Email: strings.Repeat("a", 309) + "@example.com",
	})
	if assert.Error(t, err) {
		assert.Equal(t, expected, err.Error())
	}
}

func TestCreateUserHandler_validate_Error_Emailがフォーマットエラー(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockCreateUserService(ctrl)

	handler := &createUserHandler{
		db:                db,
		createUserService: mockService,
	}
	expected := "email: must be in a valid format."
	err := handler.validate(ctx, &pb.CreateUserRequest{
		Email: "format error",
	})
	if assert.Error(t, err) {
		assert.Equal(t, expected, err.Error())
	}
}
