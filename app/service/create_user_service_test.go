package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/testutil/mock"
	"github.com/takahiroaoki/grpc-sample/app/util"
)

func Test_createUserServiceImpl_CreateUser(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockDemoRepository(ctrl)

	type args struct {
		dr *mock.MockDemoRepository
		u  entity.User
	}
	tests := []struct {
		name        string
		service     *createUserServiceImpl
		args        args
		mockFunc    func(mockRepository *mock.MockDemoRepository)
		expected    *entity.User
		isError     bool
		expectedErr util.AppError
	}{
		{
			name:    "Success",
			service: &createUserServiceImpl{},
			args: args{
				dr: mockRepository,
				u: entity.User{
					Email: "user@example.com",
				},
			},
			mockFunc: func(mockRepository *mock.MockDemoRepository) {
				mockRepository.EXPECT().CreateOneUser(entity.User{
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
			name:    "Error(service is nil)",
			service: nil,
			args: args{
				dr: mockRepository,
				u: entity.User{
					Email: "user@example.com",
				},
			},
			expected:    nil,
			isError:     true,
			expectedErr: util.NewAppErrorFromMsg("*createUserServiceImpl is nil", util.CAUSE_INTERNAL, util.LOG_LEVEL_ERROR),
		},
		{
			name:    "Error(CreateOneUser)",
			service: &createUserServiceImpl{},
			args: args{
				dr: mockRepository,
				u: entity.User{
					Email: "user@example.com",
				},
			},
			mockFunc: func(mockRepository *mock.MockDemoRepository) {
				mockRepository.EXPECT().CreateOneUser(entity.User{
					Email: "user@example.com",
				}).Return(nil, util.NewAppErrorFromMsg("err", util.CAUSE_UNDEFINED, util.LOG_LEVEL_UNDEFINED))
			},
			expected:    nil,
			isError:     true,
			expectedErr: util.NewAppErrorFromMsg("err", util.CAUSE_UNDEFINED, util.LOG_LEVEL_UNDEFINED),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFunc != nil {
				tt.mockFunc(tt.args.dr)
			}
			actual, err := tt.service.CreateUser(tt.args.dr, tt.args.u)

			assert.Equal(t, tt.expected, actual)
			if tt.isError {
				assert.Error(t, err)
				assert.True(t, err.Equals(tt.expectedErr))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
