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

func Test_createUserHandlerImpl_Invoke(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		req *CreateUserRequest
	}
	tests := []struct {
		name        string
		args        args
		mockFunc    func(mockService *mockservice.MockCreateUserService)
		expected    *CreateUserResponse
		isError     bool
		expectedErr domerr.DomErr
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				req: &CreateUserRequest{
					email: "user@example.com",
				},
			},
			mockFunc: func(mockService *mockservice.MockCreateUserService) {
				mockService.EXPECT().CreateUser(gomock.Any(), entity.User{
					Email: "user@example.com",
				}).Return(&entity.User{
					ID:    1,
					Email: "user@example.com",
				}, nil)
			},
			expected: &CreateUserResponse{
				id: "1",
			},
			isError: false,
		},
		{
			name: "Error(CreateUser)",
			args: args{
				ctx: context.Background(),
				req: &CreateUserRequest{
					email: "user@example.com",
				},
			},
			mockFunc: func(mockService *mockservice.MockCreateUserService) {
				mockService.EXPECT().CreateUser(gomock.Any(), entity.User{
					Email: "user@example.com",
				}).Return(nil, domerr.NewDomErrFromMsg("err", domerr.CAUSE_UNDEFINED, domerr.LOG_LEVEL_UNDEFINED))
			},
			expected:    nil,
			isError:     true,
			expectedErr: domerr.NewDomErrFromMsg("err", domerr.CAUSE_UNDEFINED, domerr.LOG_LEVEL_UNDEFINED),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockService := mockservice.NewMockCreateUserService(ctrl)
			if tt.mockFunc != nil {
				tt.mockFunc(mockService)
			}
			handler := &createUserHandlerImpl{
				cus: mockService,
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
