package handler

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/testutil"
)

func Test_getUserInfoHandlerImpl_Invoke(t *testing.T) {
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
		mockFunc    func(mockService *MockgetUserInfoService)
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
			mockFunc: func(mockService *MockgetUserInfoService) {
				mockService.EXPECT().GetUserByUserId(gomock.Any(), "1").Return(&entity.User{
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
			mockFunc: func(mockService *MockgetUserInfoService) {
				mockService.EXPECT().GetUserByUserId(gomock.Any(), "1").Return(nil, domerr.NewDomErrFromMsg("err", domerr.CAUSE_UNDEFINED, domerr.LOG_LEVEL_UNDEFINED))
			},
			expected:    nil,
			isError:     true,
			expectedErr: domerr.NewDomErrFromMsg("err", domerr.CAUSE_UNDEFINED, domerr.LOG_LEVEL_UNDEFINED),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := NewMockgetUserInfoService(ctrl)
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
