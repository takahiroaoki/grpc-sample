package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/testutil/mock"
)

func Test_getUserInfoServiceImpl_GetUserByUserId(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockDemoRepository(ctrl)

	type args struct {
		dr     *mock.MockDemoRepository
		userId string
	}
	tests := []struct {
		name           string
		service        *getUserInfoServiceImpl
		args           args
		mockFunc       func(mockRepository *mock.MockDemoRepository)
		expected       *entity.User
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name:    "Success",
			service: &getUserInfoServiceImpl{},
			args: args{
				dr:     mockRepository,
				userId: "1",
			},
			mockFunc: func(mockRepository *mock.MockDemoRepository) {
				mockRepository.EXPECT().SelectOneUserByUserId("1").Return(&entity.User{
					ID:    1,
					Email: "user@example.com",
				}, nil)
			},
			expected: &entity.User{
				ID:    1,
				Email: "user@example.com",
			},
			expectErr: false,
		},
		{
			name:    "Error(service is nil)",
			service: nil,
			args: args{
				dr:     mockRepository,
				userId: "1",
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "*getUserInfoServiceImpl is nil",
		},
		{
			name:    "Error(SelectOneUserByUserId)",
			service: &getUserInfoServiceImpl{},
			args: args{
				dr:     mockRepository,
				userId: "1",
			},
			mockFunc: func(mockRepository *mock.MockDemoRepository) {
				mockRepository.EXPECT().SelectOneUserByUserId("1").Return(nil, errors.New("err"))
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "err",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFunc != nil {
				tt.mockFunc(tt.args.dr)
			}
			actual, err := tt.service.GetUserByUserId(tt.args.dr, tt.args.userId)

			assert.Equal(t, tt.expected, actual)
			if tt.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErrMsg, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
