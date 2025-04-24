package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
	"github.com/takahiroaoki/grpc-sample/app/testutil"
	"github.com/takahiroaoki/grpc-sample/app/testutil/mockrepository"
)

func Test_createUserService_CreateUser(t *testing.T) {
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
		mockFunc    func(mockRepository *mockrepository.MockDemoRepository)
		expected    entity.User
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
			mockFunc: func(mockRepository *mockrepository.MockDemoRepository) {
				mockRepository.EXPECT().Transaction(gomock.Any()).DoAndReturn(func(fn func(tx repository.DemoRepository) error) error {
					return fn(mockRepository)
				})
				mockRepository.EXPECT().CreateOneUser(gomock.Any(), entity.User{
					Email: "user@example.com",
				}).Return(entity.User{
					ID:    1,
					Email: "user@example.com",
				}, nil)
			},
			expected: entity.User{
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
			mockFunc: func(mockRepository *mockrepository.MockDemoRepository) {
				mockRepository.EXPECT().Transaction(gomock.Any()).DoAndReturn(func(fn func(tx repository.DemoRepository) error) error {
					return fn(mockRepository)
				})
				mockRepository.EXPECT().CreateOneUser(gomock.Any(), entity.User{
					Email: "user@example.com",
				}).Return(entity.User{}, domerr.NewDomErrFromMsg("err", domerr.CAUSE_UNDEFINED, domerr.LOG_LEVEL_UNDEFINED))
			},
			expected:    entity.User{},
			isError:     true,
			expectedErr: domerr.NewDomErrFromMsg("createUserService.CreateUser: err", domerr.CAUSE_UNDEFINED, domerr.LOG_LEVEL_UNDEFINED),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockRepository := mockrepository.NewMockDemoRepository(ctrl)
			if tt.mockFunc != nil {
				tt.mockFunc(mockRepository)
			}
			service := &createUserService{
				dr: mockRepository,
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
