package handler

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/testutil"
	"github.com/takahiroaoki/grpc-sample/app/testutil/mockservice"
)

func Test_getUserInfoHandlerImpl_Invoke(t *testing.T) {
	t.Parallel()

	dbc, _, err := testutil.GetMockDBClient()
	assert.NoError(t, err)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mockservice.NewMockGetUserInfoService(ctrl)

	type args struct {
		ctx context.Context
		req *GetUserInfoRequest
	}
	tests := []struct {
		name        string
		handler     *getUserInfoHandlerImpl
		args        args
		mockFunc    func(mockRepository *mockservice.MockGetUserInfoService)
		expected    *GetUserInfoResponse
		isError     bool
		expectedErr domerr.DomErr
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
			mockFunc: func(mockService *mockservice.MockGetUserInfoService) {
				mockService.EXPECT().GetUserByUserId(gomock.Any(), dbc, "1").Return(&entity.User{
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
			mockFunc: func(mockService *mockservice.MockGetUserInfoService) {
				mockService.EXPECT().GetUserByUserId(gomock.Any(), dbc, "1").Return(nil, domerr.NewDomErrFromMsg("err", domerr.CAUSE_UNDEFINED, domerr.LOG_LEVEL_UNDEFINED))
			},
			expected:    nil,
			isError:     true,
			expectedErr: domerr.NewDomErrFromMsg("err", domerr.CAUSE_UNDEFINED, domerr.LOG_LEVEL_UNDEFINED),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFunc != nil {
				tt.mockFunc(mockService)
			}
			actual, err := tt.handler.Invoke(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.expected, actual)
			if tt.isError {
				assert.Error(t, err)
				assert.True(t, testutil.SameDomainErrors(err, tt.expectedErr))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
