package repository

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/go-env/app/entity"
	"github.com/takahiroaoki/go-env/app/testutil"
	"gorm.io/gorm"
)

func Test_userRepositoryImpl_CreateOneUser(t *testing.T) {
	db, _ := testutil.GetDatabase()

	type args struct {
		db *gorm.DB
		u  entity.User
	}
	tests := []struct {
		name           string
		args           args
		preprocess     func(db *gorm.DB, userRepository UserRepository)
		expected       *entity.User
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name: "Success",
			args: args{
				db: db,
				u: entity.User{
					ID:    1,
					Email: "CreateOneUser@example.com",
				},
			},
			preprocess: func(db *gorm.DB, userRepository UserRepository) {
				if err := testutil.CleanDB(db); err != nil {
					t.FailNow()
				}
			},
			expected: &entity.User{
				ID:    1,
				Email: "CreateOneUser@example.com",
			},
			expectErr: false,
		},
		{
			name: "Success(Email has 320 length)",
			args: args{
				db: db,
				u: entity.User{
					ID:    1,
					Email: strings.Repeat("a", 308) + "@example.com",
				},
			},
			preprocess: func(db *gorm.DB, userRepository UserRepository) {
				if err := testutil.CleanDB(db); err != nil {
					t.FailNow()
				}
			},
			expected: &entity.User{
				ID:    1,
				Email: strings.Repeat("a", 308) + "@example.com",
			},
			expectErr: false,
		},
		{
			name: "Error(Email has 321 length)",
			args: args{
				db: db,
				u: entity.User{
					Email: strings.Repeat("a", 309) + "@example.com",
				},
			},
			preprocess: func(db *gorm.DB, userRepository UserRepository) {
				if err := testutil.CleanDB(db); err != nil {
					t.FailNow()
				}
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "Error 1406 (22001): Data too long for column 'email' at row 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &userRepositoryImpl{}
			if tt.preprocess != nil {
				tt.preprocess(db, r)
			}
			actual, err := r.CreateOneUser(tt.args.db, tt.args.u)

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

func Test_userRepositoryImpl_SelectOneUserByUserId(t *testing.T) {
	db, _ := testutil.GetDatabase()

	type args struct {
		db     *gorm.DB
		userId string
	}
	tests := []struct {
		name           string
		args           args
		preprocess     func(db *gorm.DB, userRepository UserRepository)
		expected       *entity.User
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name: "Success",
			args: args{
				db:     db,
				userId: "1",
			},
			preprocess: func(db *gorm.DB, userRepository UserRepository) {
				if err := testutil.CleanDB(db); err != nil {
					t.FailNow()
				}
				if _, err := userRepository.CreateOneUser(db, entity.User{
					ID:    1,
					Email: "repository@example.com",
				}); err != nil {
					t.FailNow()
				}
			},
			expected: &entity.User{
				ID:    1,
				Email: "repository@example.com",
			},
			expectErr: false,
		},
		{
			name: "Error",
			args: args{
				db:     db,
				userId: "no existing value",
			},
			expected:       nil,
			expectErr:      true,
			expectedErrMsg: "record not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &userRepositoryImpl{}
			if tt.preprocess != nil {
				tt.preprocess(db, r)
			}
			actual, err := r.SelectOneUserByUserId(tt.args.db, tt.args.userId)

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
