package handler

import (
	"context"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/pb"
	"github.com/takahiroaoki/grpc-sample/app/repository"
	"github.com/takahiroaoki/grpc-sample/app/testutil"
	"github.com/takahiroaoki/grpc-sample/app/testutil/mock"
	"github.com/takahiroaoki/grpc-sample/app/util"
)

func Test_createUserHandlerImpl_Execute(t *testing.T) {
	t.Parallel()

	dbc, sqlMock, err := testutil.GetMockDBClient()
	assert.NoError(t, err)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockCreateUserService(ctrl)

	type fields struct {
		dr  repository.DemoRepository
		cus *mock.MockCreateUserService
	}
	type args struct {
		ctx context.Context
		req *pb.CreateUserRequest
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		mockFunc       func(sqlMock sqlmock.Sqlmock, mockRepository *mock.MockCreateUserService)
		expected       *pb.CreateUserResponse
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name: "Success",
			fields: fields{
				dr:  dbc,
				cus: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.CreateUserRequest{
					Email: "user@example.com",
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
			expected: &pb.CreateUserResponse{
				Id: "1",
			},
			expectErr: false,
		},
		{
			name: "Error(validation)",
			fields: fields{
				dr:  dbc,
				cus: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.CreateUserRequest{
					Email: "invalid value",
				},
			},
			mockFunc: func(sqlMock sqlmock.Sqlmock, mockService *mock.MockCreateUserService) {
				mockService.EXPECT().CreateUser(gomock.Any(), gomock.Any()).MaxTimes(0)
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "email: must be in a valid format.",
		},
		{
			name: "Error(createUserService.CreateUser)",
			fields: fields{
				dr:  dbc,
				cus: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.CreateUserRequest{
					Email: "user@example.com",
				},
			},
			mockFunc: func(sqlMock sqlmock.Sqlmock, mockService *mock.MockCreateUserService) {
				sqlMock.ExpectBegin()
				mockService.EXPECT().CreateUser(gomock.Any(), entity.User{
					Email: "user@example.com",
				}).Return(nil, util.NewError("err"))
				sqlMock.ExpectRollback()
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "err",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &createUserHandlerImpl{
				dr:  tt.fields.dr,
				cus: tt.fields.cus,
			}
			tt.mockFunc(sqlMock, tt.fields.cus)
			actual, err := h.Execute(tt.args.ctx, tt.args.req)

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

	type fields struct {
		dr  repository.DemoRepository
		cus *mock.MockCreateUserService
	}
	type args struct {
		ctx context.Context
		req *pb.CreateUserRequest
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		expected       error
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name: "Success",
			fields: fields{
				dr:  dbc,
				cus: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.CreateUserRequest{
					Email: "user@example.com",
				},
			},
			expectErr: false,
		},
		{
			name: "Success(Email right boundary safe)",
			fields: fields{
				dr:  dbc,
				cus: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.CreateUserRequest{
					Email: strings.Repeat("a", 308) + "@example.com",
				},
			},
			expectErr: false,
		},
		{
			name: "Error(Email right boundary over)",
			fields: fields{
				dr:  dbc,
				cus: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.CreateUserRequest{
					Email: strings.Repeat("a", 309) + "@example.com",
				},
			},
			expectErr:      true,
			expectedErrMsg: "email: the length must be between 1 and 320.",
		},
		{
			name: "Error(Email is nil)",
			fields: fields{
				dr:  dbc,
				cus: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.CreateUserRequest{},
			},
			expectErr:      true,
			expectedErrMsg: "email: cannot be blank.",
		},
		{
			name: "Error(Email is empty)",
			fields: fields{
				dr:  dbc,
				cus: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.CreateUserRequest{
					Email: "",
				},
			},
			expectErr:      true,
			expectedErrMsg: "email: cannot be blank.",
		},
		{
			name: "Error(Email is in an invalid format)",
			fields: fields{
				dr:  dbc,
				cus: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.CreateUserRequest{
					Email: "invalid format",
				},
			},
			expectErr:      true,
			expectedErrMsg: "email: must be in a valid format.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &createUserHandlerImpl{
				dr:  tt.fields.dr,
				cus: tt.fields.cus,
			}
			err := h.validate(tt.args.ctx, tt.args.req)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErrMsg, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
