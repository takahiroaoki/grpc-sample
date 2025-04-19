package handler

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/testutil"
	"github.com/takahiroaoki/grpc-sample/app/testutil/mockservice"
)

func Test_createUserHandlerImpl_Invoke(t *testing.T) {
	t.Parallel()

	dbc, sqlMock, err := testutil.GetMockDBClient()
	assert.NoError(t, err)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mockservice.NewMockCreateUserService(ctrl)

	type args struct {
		ctx context.Context
		req *CreateUserRequest
	}
	tests := []struct {
		name        string
		handler     *createUserHandlerImpl
		args        args
		mockFunc    func(sqlMock sqlmock.Sqlmock, mockRepository *mockservice.MockCreateUserService)
		expected    *CreateUserResponse
		isError     bool
		expectedErr domerr.DomErr
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
			mockFunc: func(sqlMock sqlmock.Sqlmock, mockService *mockservice.MockCreateUserService) {
				sqlMock.ExpectBegin()
				mockService.EXPECT().CreateUser(gomock.Any(), gomock.Any(), entity.User{
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
			isError: false,
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
			mockFunc: func(sqlMock sqlmock.Sqlmock, mockService *mockservice.MockCreateUserService) {
				sqlMock.ExpectBegin()
				mockService.EXPECT().CreateUser(gomock.Any(), gomock.Any(), entity.User{
					Email: "user@example.com",
				}).Return(nil, domerr.NewDomErrFromMsg("err", domerr.CAUSE_UNDEFINED, domerr.LOG_LEVEL_UNDEFINED))
				sqlMock.ExpectRollback()
			},
			expected:    nil,
			isError:     true,
			expectedErr: domerr.NewDomErrFromMsg("err", domerr.CAUSE_UNDEFINED, domerr.LOG_LEVEL_UNDEFINED),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFunc != nil {
				tt.mockFunc(sqlMock, mockService)
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
