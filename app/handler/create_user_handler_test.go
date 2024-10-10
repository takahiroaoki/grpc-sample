package handler

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/testutil"
	"github.com/takahiroaoki/grpc-sample/app/testutil/mock"
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
		name           string
		handler        *createUserHandlerImpl
		args           args
		mockFunc       func(sqlMock sqlmock.Sqlmock, mockRepository *mock.MockCreateUserService)
		expected       *CreateUserResponse
		expectErr      bool
		expectedErrMsg string
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
			expectErr: false,
		},
		{
			name:    "Error(handler is nil)",
			handler: nil,
			args: args{
				ctx: context.Background(),
				req: &CreateUserRequest{
					email: "invalid value",
				},
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "*createUserHandlerImpl is nil",
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
				}).Return(nil, errors.New("err"))
				sqlMock.ExpectRollback()
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "err",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFunc != nil {
				tt.mockFunc(sqlMock, mockService)
			}
			actual, err := tt.handler.process(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.expected, actual)
			if tt.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErrMsg, err.Error())
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
		name           string
		handler        *createUserHandlerImpl
		args           args
		expected       error
		expectErr      bool
		expectedErrMsg string
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
			expectErr: false,
		},
		{
			name:    "Error(handler is nil)",
			handler: nil,
			args: args{
				ctx: context.Background(),
				req: &CreateUserRequest{
					email: strings.Repeat("a", 309) + "@example.com",
				},
			},
			expectErr:      true,
			expectedErrMsg: "*createUserHandlerImpl is nil",
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
			expectErr: false,
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
			expectErr:      true,
			expectedErrMsg: "email: the length must be between 1 and 320.",
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
			expectErr:      true,
			expectedErrMsg: "email: cannot be blank.",
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
			expectErr:      true,
			expectedErrMsg: "email: cannot be blank.",
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
			expectErr:      true,
			expectedErrMsg: "email: must be in a valid format.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.handler.validate(tt.args.ctx, tt.args.req)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErrMsg, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
