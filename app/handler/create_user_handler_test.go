package handler

import (
	"context"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/testutil"
	"github.com/takahiroaoki/grpc-sample/app/testutil/mock"
	"github.com/takahiroaoki/grpc-sample/app/util"
)

func Test_createUserHandlerImpl_process(t *testing.T) {
	t.Parallel()

	dbc, sqlMock, err := testutil.GetMockDBClient()
	assert.NoError(t, err)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockCreateUserService(ctrl)

	type args struct {
		ctx context.Context
		req *CreateUserRequest
	}
	tests := []struct {
		name        string
		handler     *createUserHandlerImpl
		args        args
		mockFunc    func(sqlMock sqlmock.Sqlmock, mockRepository *mock.MockCreateUserService)
		expected    *CreateUserResponse
		isError     bool
		expectedErr util.AppError
	}{
		{
			name: "Success",
			handler: &createUserHandlerImpl{
				dr:  dbc,
				cus: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &CreateUserRequest{
					email: "user@example.com",
				},
			},
			mockFunc: func(sqlMock sqlmock.Sqlmock, mockService *mock.MockCreateUserService) {
				sqlMock.ExpectBegin()
				mockService.EXPECT().CreateUser(gomock.Any(), entity.User{
					Email: "user@example.com",
				}).Return(&entity.User{
					ID:    1,
					Email: "user@example.com",
				}, nil)
				sqlMock.ExpectCommit()
			},
			expected: &CreateUserResponse{
				id: "1",
			},
			isError: false,
		},
		{
			name:    "Error(handler is nil)",
			handler: nil,
			args: args{
				ctx: context.Background(),
				req: &CreateUserRequest{
					email: "user@example.com",
				},
			},
			expected:    nil,
			isError:     true,
			expectedErr: util.NewAppErrorFromMsg("*createUserHandlerImpl is nil", util.CAUSE_INTERNAL, util.LOG_LEVEL_ERROR),
		},
		{
			name: "Error(CreateUser)",
			handler: &createUserHandlerImpl{
				dr:  dbc,
				cus: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &CreateUserRequest{
					email: "user@example.com",
				},
			},
			mockFunc: func(sqlMock sqlmock.Sqlmock, mockService *mock.MockCreateUserService) {
				sqlMock.ExpectBegin()
				mockService.EXPECT().CreateUser(gomock.Any(), entity.User{
					Email: "user@example.com",
				}).Return(nil, util.NewAppErrorFromMsg("err", util.CAUSE_UNDEFINED, util.LOG_LEVEL_UNDEFINED))
				sqlMock.ExpectRollback()
			},
			expected:    nil,
			isError:     true,
			expectedErr: util.NewAppErrorFromMsg("err", util.CAUSE_UNDEFINED, util.LOG_LEVEL_UNDEFINED),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFunc != nil {
				tt.mockFunc(sqlMock, mockService)
			}
			actual, err := tt.handler.process(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.expected, actual)
			if tt.isError {
				assert.Error(t, err)
				assert.True(t, err.Is(tt.expectedErr))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_createUserHandlerImpl_validate(t *testing.T) {
	t.Parallel()

	dbc, _, err := testutil.GetMockDBClient()
	assert.NoError(t, err)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockCreateUserService(ctrl)

	type args struct {
		ctx context.Context
		req *CreateUserRequest
	}
	tests := []struct {
		name        string
		handler     *createUserHandlerImpl
		args        args
		expected    error
		isError     bool
		expectedErr util.AppError
	}{
		{
			name: "Success",
			handler: &createUserHandlerImpl{
				dr:  dbc,
				cus: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &CreateUserRequest{
					email: "user@example.com",
				},
			},
			isError: false,
		},
		{
			name:    "Error(handler is nil)",
			handler: nil,
			args: args{
				ctx: context.Background(),
				req: &CreateUserRequest{
					email: "user@example.com",
				},
			},
			isError:     true,
			expectedErr: util.NewAppErrorFromMsg("*createUserHandlerImpl is nil", util.CAUSE_INTERNAL, util.LOG_LEVEL_ERROR),
		},
		{
			name: "Success(Email right boundary safe)",
			handler: &createUserHandlerImpl{
				dr:  dbc,
				cus: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &CreateUserRequest{
					email: strings.Repeat("a", 308) + "@example.com",
				},
			},
			isError: false,
		},
		{
			name: "Error(Email right boundary over)",
			handler: &createUserHandlerImpl{
				dr:  dbc,
				cus: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &CreateUserRequest{
					email: strings.Repeat("a", 309) + "@example.com",
				},
			},
			isError:     true,
			expectedErr: util.NewAppErrorFromMsg("email: the length must be between 1 and 320.", util.CAUSE_INVALID_ARGUMENT, util.LOG_LEVEL_INFO),
		},
		{
			name: "Error(Email is nil)",
			handler: &createUserHandlerImpl{
				dr:  dbc,
				cus: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &CreateUserRequest{},
			},
			isError:     true,
			expectedErr: util.NewAppErrorFromMsg("email: cannot be blank.", util.CAUSE_INVALID_ARGUMENT, util.LOG_LEVEL_INFO),
		},
		{
			name: "Error(Email is empty)",
			handler: &createUserHandlerImpl{
				dr:  dbc,
				cus: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &CreateUserRequest{
					email: "",
				},
			},
			isError:     true,
			expectedErr: util.NewAppErrorFromMsg("email: cannot be blank.", util.CAUSE_INVALID_ARGUMENT, util.LOG_LEVEL_INFO),
		},
		{
			name: "Error(Email is in an invalid format)",
			handler: &createUserHandlerImpl{
				dr:  dbc,
				cus: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &CreateUserRequest{
					email: "invalid format",
				},
			},
			isError:     true,
			expectedErr: util.NewAppErrorFromMsg("email: must be in a valid format.", util.CAUSE_INVALID_ARGUMENT, util.LOG_LEVEL_INFO),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.handler.validate(tt.args.ctx, tt.args.req)

			if tt.isError {
				assert.Error(t, err)
				assert.True(t, err.Is(tt.expectedErr))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
