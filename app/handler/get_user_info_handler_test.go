package handler

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/pb"
	"github.com/takahiroaoki/grpc-sample/app/testutil"
	"github.com/takahiroaoki/grpc-sample/app/testutil/mock"
	"github.com/takahiroaoki/grpc-sample/app/util"
	"gorm.io/gorm"
)

func Test_getUserInfoHandlerImpl_execute(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockGetUserInfoService(ctrl)

	type fields struct {
		db                 *gorm.DB
		getUserInfoService *mock.MockGetUserInfoService
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
				db:                 db,
				getUserInfoService: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.GetUserInfoRequest{
					Id: "1",
				},
			},
			mockFunc: func(mockService *mock.MockGetUserInfoService) {
				mockService.EXPECT().GetUserByUserId(db, "1").Return(&entity.User{
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
				db:                 db,
				getUserInfoService: mockService,
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
				db:                 db,
				getUserInfoService: mockService,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.GetUserInfoRequest{
					Id: "1",
				},
			},
			mockFunc: func(mockService *mock.MockGetUserInfoService) {
				mockService.EXPECT().GetUserByUserId(db, "1").Return(nil, util.NewError("err"))
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "err",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			h := &getUserInfoHandlerImpl{
				db:                 tt.fields.db,
				getUserInfoService: tt.fields.getUserInfoService,
			}
			tt.mockFunc(tt.fields.getUserInfoService)
			actual, err := h.execute(tt.args.ctx, tt.args.req)

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

	db, _ := testutil.GetDatabase()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockGetUserInfoService(ctrl)

	type fields struct {
		db                 *gorm.DB
		getUserInfoService *mock.MockGetUserInfoService
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
				db:                 db,
				getUserInfoService: mockService,
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
				db:                 db,
				getUserInfoService: mockService,
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
				db:                 db,
				getUserInfoService: mockService,
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
				db:                 db,
				getUserInfoService: mockService,
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
			t.Parallel()
			h := &getUserInfoHandlerImpl{
				db:                 tt.fields.db,
				getUserInfoService: tt.fields.getUserInfoService,
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
