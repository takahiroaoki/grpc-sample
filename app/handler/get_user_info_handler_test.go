package handler

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
	"github.com/takahiroaoki/grpc-sample/app/infra/pb"
	"github.com/takahiroaoki/grpc-sample/app/testutil"
	"github.com/takahiroaoki/grpc-sample/app/testutil/mock"
	"github.com/takahiroaoki/grpc-sample/app/util"
)

func Test_getUserInfoHandlerImpl_Execute(t *testing.T) {
	t.Parallel()

	dbc, _, err := testutil.GetMockDBClient()
	assert.NoError(t, err)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockGetUserInfoService(ctrl)

	type fields struct {
		dr   repository.DemoRepository
		guis *mock.MockGetUserInfoService
	}
	type args struct {
		ctx context.Context
		req *pb.GetUserInfoRequest
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		mockFunc       func(mockRepository *mock.MockGetUserInfoService)
		expected       *pb.GetUserInfoResponse
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name: "Success",
			fields: fields{
				dr:   dbc,
				guis: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.GetUserInfoRequest{
					Id: "1",
				},
			},
			mockFunc: func(mockService *mock.MockGetUserInfoService) {
				mockService.EXPECT().GetUserByUserId(dbc, "1").Return(&entity.User{
					ID:    1,
					Email: "user@example.com",
				}, nil)
			},
			expected: &pb.GetUserInfoResponse{
				Id:    "1",
				Email: "user@example.com",
			},
			expectErr: false,
		},
		{
			name: "Error(validation)",
			fields: fields{
				dr:   dbc,
				guis: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.GetUserInfoRequest{
					Id: "invalid value",
				},
			},
			mockFunc: func(mockService *mock.MockGetUserInfoService) {
				mockService.EXPECT().GetUserByUserId(gomock.Any(), gomock.Any()).MaxTimes(0)
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "id: must contain digits only.",
		},
		{
			name: "Error(getUserInfoService.GetUserByUserId)",
			fields: fields{
				dr:   dbc,
				guis: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.GetUserInfoRequest{
					Id: "1",
				},
			},
			mockFunc: func(mockService *mock.MockGetUserInfoService) {
				mockService.EXPECT().GetUserByUserId(dbc, "1").Return(nil, util.NewError("err"))
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "err",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &getUserInfoHandlerImpl{
				dr:   tt.fields.dr,
				guis: tt.fields.guis,
			}
			tt.mockFunc(tt.fields.guis)
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

func Test_getUserInfoHandlerImpl_validate(t *testing.T) {
	t.Parallel()

	dbc, _, err := testutil.GetMockDBClient()
	assert.NoError(t, err)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockGetUserInfoService(ctrl)

	type fields struct {
		dr   repository.DemoRepository
		guis *mock.MockGetUserInfoService
	}
	type args struct {
		ctx context.Context
		req *pb.GetUserInfoRequest
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
				dr:   dbc,
				guis: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.GetUserInfoRequest{
					Id: "12345",
				},
			},
			expected:  nil,
			expectErr: false,
		},
		{
			name: "Error(Id is nil)",
			fields: fields{
				dr:   dbc,
				guis: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.GetUserInfoRequest{},
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "id: cannot be blank.",
		},
		{
			name: "Error(Id is empty)",
			fields: fields{
				dr:   dbc,
				guis: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.GetUserInfoRequest{
					Id: "",
				},
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "id: cannot be blank.",
		},
		{
			name: "Error(Id contains invalid characters)",
			fields: fields{
				dr:   dbc,
				guis: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.GetUserInfoRequest{
					Id: "abc",
				},
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "id: must contain digits only.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &getUserInfoHandlerImpl{
				dr:   tt.fields.dr,
				guis: tt.fields.guis,
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
