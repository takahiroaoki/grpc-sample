package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/infra"
	"github.com/takahiroaoki/grpc-sample/app/testutil"
	"github.com/takahiroaoki/grpc-sample/app/testutil/mock"
	"github.com/takahiroaoki/grpc-sample/app/util"
)

func Test_getUserInfoServiceImpl_GetUserByUserId(t *testing.T) {
	t.Parallel()

	dbw, _, err := testutil.GetTestDBWrapper()
	assert.NoError(t, err)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockDemoRepository(ctrl)

	type fields struct {
		demoRepository *mock.MockDemoRepository
	}
	type args struct {
		dbw    infra.DBWrapper
		userId string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		mockFunc       func(mockRepository *mock.MockDemoRepository)
		expected       *entity.User
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name: "Success",
			fields: fields{
				demoRepository: mockRepository,
			},
			args: args{
				dbw:    dbw,
				userId: "1",
			},
			mockFunc: func(mockRepository *mock.MockDemoRepository) {
				mockRepository.EXPECT().SelectOneUserByUserId(dbw, "1").Return(&entity.User{
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
			name: "Error",
			fields: fields{
				demoRepository: mockRepository,
			},
			args: args{
				dbw:    dbw,
				userId: "1",
			},
			mockFunc: func(mockRepository *mock.MockDemoRepository) {
				mockRepository.EXPECT().SelectOneUserByUserId(dbw, "1").Return(nil, util.NewError("err"))
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "err",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := &getUserInfoServiceImpl{
				demoRepository: tt.fields.demoRepository,
			}
			tt.mockFunc(tt.fields.demoRepository)
			actual, err := s.GetUserByUserId(tt.args.dbw, tt.args.userId)

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
