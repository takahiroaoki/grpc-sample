package handler

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/testutil"
	"github.com/takahiroaoki/grpc-sample/app/testutil/mock"
	"github.com/takahiroaoki/grpc-sample/app/util"
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
		name        string
		handler     *getUserInfoHandlerImpl
		args        args
		mockFunc    func(mockRepository *mock.MockGetUserInfoService)
		expected    *GetUserInfoResponse
		isError     bool
		expectedErr util.AppError
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
			isError: false,
		},
		{
			name:    "Error(handler is nil)",
			handler: nil,
			args: args{
				ctx: context.Background(),
				req: &GetUserInfoRequest{
					id: "1",
				},
			},
			expected:    nil,
			isError:     true,
			expectedErr: util.NewAppErrorFromMsg("*getUserInfoHandlerImpl is nil", util.CAUSE_INTERNAL, util.LOG_LEVEL_ERROR),
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
				mockService.EXPECT().GetUserByUserId(dbc, "1").Return(nil, util.NewAppErrorFromMsg("err", util.CAUSE_UNDEFINED, util.LOG_LEVEL_UNDEFINED))
			},
			expected:    nil,
			isError:     true,
			expectedErr: util.NewAppErrorFromMsg("err", util.CAUSE_UNDEFINED, util.LOG_LEVEL_UNDEFINED),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFunc != nil {
				tt.mockFunc(mockService)
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
		name        string
		handler     *getUserInfoHandlerImpl
		args        args
		expected    error
		isError     bool
		expectedErr util.AppError
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
			expected: nil,
			isError:  false,
		},
		{
			name:    "Error(handler is nil)",
			handler: nil,
			args: args{
				ctx: context.Background(),
				req: &GetUserInfoRequest{
					id: "12345",
				},
			},
			expected:    nil,
			isError:     true,
			expectedErr: util.NewAppErrorFromMsg("*getUserInfoHandlerImpl is nil", util.CAUSE_INTERNAL, util.LOG_LEVEL_ERROR),
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
			expected:    nil,
			isError:     true,
			expectedErr: util.NewAppErrorFromMsg("id: cannot be blank.", util.CAUSE_INVALID_ARGUMENT, util.LOG_LEVEL_INFO),
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
			expected:    nil,
			isError:     true,
			expectedErr: util.NewAppErrorFromMsg("id: cannot be blank.", util.CAUSE_INVALID_ARGUMENT, util.LOG_LEVEL_INFO),
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
			expected:    nil,
			isError:     true,
			expectedErr: util.NewAppErrorFromMsg("id: must contain digits only.", util.CAUSE_INVALID_ARGUMENT, util.LOG_LEVEL_INFO),
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
