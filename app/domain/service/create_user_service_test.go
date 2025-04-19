package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/testutil"
)

func Test_createUserServiceImpl_CreateUser(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		u   entity.User
	}
	tests := []struct {
		name        string
		args        args
		mockFunc    func(mockRepository *MockcreateUserRepository)
		expected    *entity.User
		isError     bool
		expectedErr domerr.DomErr
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				u: entity.User{
					Email: "user@example.com",
				},
			},
			mockFunc: func(mockRepository *MockcreateUserRepository) {
				mockRepository.EXPECT().CreateOneUser(gomock.Any(), entity.User{
					Email: "user@example.com",
				}).Return(&entity.User{
					ID:    1,
					Email: "user@example.com",
				}, nil)
			},
			expected: &entity.User{
				ID:    1,
				Email: "user@example.com",
			},
			isError: false,
		},
		{
			name: "Error(CreateOneUser)",
			args: args{
				ctx: context.Background(),
				u: entity.User{
					Email: "user@example.com",
				},
			},
			mockFunc: func(mockRepository *MockcreateUserRepository) {
				mockRepository.EXPECT().CreateOneUser(gomock.Any(), entity.User{
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
			mockRepository := NewMockcreateUserRepository(ctrl)
			if tt.mockFunc != nil {
				tt.mockFunc(mockRepository)
			}
			service := &createUserService{
				cur: mockRepository,
			}
			actual, err := service.CreateUser(tt.args.ctx, tt.args.u)

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
