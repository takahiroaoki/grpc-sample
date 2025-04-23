package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/testutil"
	"github.com/takahiroaoki/grpc-sample/app/testutil/mockrepository"
)

func Test_getUserInfoService_GetUserByUserId(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx    context.Context
		userId string
	}
	tests := []struct {
		name        string
		args        args
		mockFunc    func(mockRepository *mockrepository.MockDemoRepository)
		expected    *entity.User
		isError     bool
		expectedErr domerr.DomErr
	}{
		{
			name: "Success",
			args: args{
				ctx:    context.Background(),
				userId: "1",
			},
			mockFunc: func(mockRepository *mockrepository.MockDemoRepository) {
				mockRepository.EXPECT().SelectOneUserByUserId(gomock.Any(), "1").Return(&entity.User{
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
			name: "Error(SelectOneUserByUserId)",
			args: args{
				ctx:    context.Background(),
				userId: "1",
			},
			mockFunc: func(mockRepository *mockrepository.MockDemoRepository) {
				mockRepository.EXPECT().SelectOneUserByUserId(gomock.Any(), "1").Return(nil, domerr.NewDomErrFromMsg("err", domerr.CAUSE_UNDEFINED, domerr.LOG_LEVEL_UNDEFINED))
			},
			expected:    nil,
			isError:     true,
			expectedErr: domerr.NewDomErrFromMsg("err", domerr.CAUSE_UNDEFINED, domerr.LOG_LEVEL_UNDEFINED),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockrepository.NewMockDemoRepository(ctrl)
			if tt.mockFunc != nil {
				tt.mockFunc(mockRepository)
			}
			service := &getUserInfoService{
				dr: mockRepository,
			}
			actual, err := service.GetUserByUserId(tt.args.ctx, tt.args.userId)

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
