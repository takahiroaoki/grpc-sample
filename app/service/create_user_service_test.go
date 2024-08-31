package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/testutil"
	"github.com/takahiroaoki/grpc-sample/app/testutil/mock"
	"github.com/takahiroaoki/grpc-sample/app/util"
	"gorm.io/gorm"
)

func Test_createUserServiceImpl_CreateUser(t *testing.T) {
	t.Parallel()

	db, _, err := testutil.GetTestDB()
	assert.NoError(t, err)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockUserRepository(ctrl)

	type fields struct {
		userRepository *mock.MockUserRepository
	}
	type args struct {
		db *gorm.DB
		u  entity.User
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		mockFunc       func(mockRepository *mock.MockUserRepository)
		expected       *entity.User
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name: "Success",
			fields: fields{
				userRepository: mockRepository,
			},
			args: args{
				db: db,
				u: entity.User{
					Email: "user@example.com",
				},
			},
			mockFunc: func(mockRepository *mock.MockUserRepository) {
				mockRepository.EXPECT().CreateOneUser(db, entity.User{
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
			expectErr: false,
		},
		{
			name: "Error",
			fields: fields{
				userRepository: mockRepository,
			},
			args: args{
				db: db,
				u: entity.User{
					Email: "user@example.com",
				},
			},
			mockFunc: func(mockRepository *mock.MockUserRepository) {
				mockRepository.EXPECT().CreateOneUser(db, entity.User{
					Email: "user@example.com",
				}).Return(nil, util.NewError("err"))
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "err",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := &createUserServiceImpl{
				userRepository: tt.fields.userRepository,
			}
			tt.mockFunc(tt.fields.userRepository)
			actual, err := s.CreateUser(tt.args.db, tt.args.u)

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
