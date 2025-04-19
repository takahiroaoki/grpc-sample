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

func Test_getUserInfoServiceImpl_GetUserByUserId(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mockrepository.NewMockDemoRepository(ctrl)

	type args struct {
		ctx    context.Context
		dr     *mockrepository.MockDemoRepository
		userId string
	}
	tests := []struct {
		name        string
		service     *getUserInfoService
		args        args
		mockFunc    func(mockRepository *mockrepository.MockDemoRepository)
		expected    *entity.User
		isError     bool
		expectedErr domerr.DomErr
	}{
		{
			name:    "Success",
			service: &getUserInfoService{},
			args: args{
				ctx:    context.Background(),
				dr:     mockRepository,
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
			name:    "Error(SelectOneUserByUserId)",
			service: &getUserInfoService{},
			args: args{
				ctx:    context.Background(),
				dr:     mockRepository,
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
			if tt.mockFunc != nil {
				tt.mockFunc(tt.args.dr)
			}
			actual, err := tt.service.GetUserByUserId(tt.args.ctx, tt.args.dr, tt.args.userId)

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
