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

func Test_getUserInfoHandler_Invoke(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		req *GetUserInfoRequest
	}
	tests := []struct {
		name        string
		args        args
		mockFunc    func(mockRepository *mockservice.MockGetUserInfoService)
		expected    *GetUserInfoResponse
		isError     bool
		expectedErr domerr.DomErr
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				req: &GetUserInfoRequest{
					id: "1",
				},
			},
			mockFunc: func(mockService *mockservice.MockGetUserInfoService) {
				mockService.EXPECT().GetUserByUserId(gomock.Any(), "1").Return(entity.User{
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
			args: args{
				ctx: context.Background(),
				req: &GetUserInfoRequest{
					id: "1",
				},
			},
			mockFunc: func(mockService *mockservice.MockGetUserInfoService) {
				mockService.EXPECT().GetUserByUserId(gomock.Any(), "1").Return(entity.User{}, domerr.NewDomErrFromMsg("err", domerr.CAUSE_UNDEFINED, domerr.LOG_LEVEL_UNDEFINED))
			},
			expected:    nil,
			isError:     true,
			expectedErr: domerr.NewDomErrFromMsg("getUserInfoHandler.Invoke: err", domerr.CAUSE_UNDEFINED, domerr.LOG_LEVEL_UNDEFINED),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockService := mockservice.NewMockGetUserInfoService(ctrl)
			if tt.mockFunc != nil {
				tt.mockFunc(mockService)
			}
			handler := &getUserInfoHandler{
				guis: mockService,
			}
			actual, err := handler.Invoke(tt.args.ctx, tt.args.req)

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
