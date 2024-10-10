package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/testutil"
	"github.com/takahiroaoki/grpc-sample/app/testutil/mock"
)

func Test_getUserInfoHandlerImpl_process(t *testing.T) {
	t.Parallel()

	dbc, _, err := testutil.GetMockDBClient()
	assert.NoError(t, err)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockGetUserInfoService(ctrl)

	type args struct {
		ctx context.Context
		req *GetUserInfoRequest
	}
	tests := []struct {
		name           string
		handler        *getUserInfoHandlerImpl
		args           args
		mockFunc       func(mockRepository *mock.MockGetUserInfoService)
		expected       *GetUserInfoResponse
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name: "Success",
			handler: &getUserInfoHandlerImpl{
				dr:   dbc,
				guis: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &GetUserInfoRequest{
					id: "1",
				},
			},
			mockFunc: func(mockService *mock.MockGetUserInfoService) {
				mockService.EXPECT().GetUserByUserId(dbc, "1").Return(&entity.User{
					ID:    1,
					Email: "user@example.com",
				}, nil)
			},
			expected: &GetUserInfoResponse{
				id:    "1",
				email: "user@example.com",
			},
			expectErr: false,
		},
		{
			name:    "Error(handler is nil)",
			handler: nil,
			args: args{
				ctx: context.Background(),
				req: &GetUserInfoRequest{
					id: "invalid value",
				},
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "*getUserInfoHandlerImpl is nil",
		},
		{
			name: "Error(GetUserByUserId)",
			handler: &getUserInfoHandlerImpl{
				dr:   dbc,
				guis: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &GetUserInfoRequest{
					id: "1",
				},
			},
			mockFunc: func(mockService *mock.MockGetUserInfoService) {
				mockService.EXPECT().GetUserByUserId(dbc, "1").Return(nil, errors.New("err"))
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "err",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFunc != nil {
				tt.mockFunc(mockService)
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

func Test_getUserInfoHandlerImpl_validate(t *testing.T) {
	t.Parallel()

	dbc, _, err := testutil.GetMockDBClient()
	assert.NoError(t, err)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockGetUserInfoService(ctrl)

	type args struct {
		ctx context.Context
		req *GetUserInfoRequest
	}
	tests := []struct {
		name           string
		handler        *getUserInfoHandlerImpl
		args           args
		expected       error
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name: "Success",
			handler: &getUserInfoHandlerImpl{
				dr:   dbc,
				guis: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &GetUserInfoRequest{
					id: "12345",
				},
			},
			expected:  nil,
			expectErr: false,
		},
		{
			name:    "Error(handler is nil)",
			handler: nil,
			args: args{
				ctx: context.Background(),
				req: &GetUserInfoRequest{},
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "*getUserInfoHandlerImpl is nil",
		},
		{
			name: "Error(Id is nil)",
			handler: &getUserInfoHandlerImpl{
				dr:   dbc,
				guis: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &GetUserInfoRequest{},
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "id: cannot be blank.",
		},
		{
			name: "Error(Id is empty)",
			handler: &getUserInfoHandlerImpl{
				dr:   dbc,
				guis: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &GetUserInfoRequest{
					id: "",
				},
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "id: cannot be blank.",
		},
		{
			name: "Error(Id contains invalid characters)",
			handler: &getUserInfoHandlerImpl{
				dr:   dbc,
				guis: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &GetUserInfoRequest{
					id: "abc",
				},
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "id: must contain digits only.",
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
