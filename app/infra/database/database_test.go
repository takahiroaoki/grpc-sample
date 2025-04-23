package database

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_DBClient_Transaction(t *testing.T) {
	dbc, sqlMock, err := GetMockDBClient()
	assert.NoError(t, err)

	type args struct {
		fn func(tx repository.DemoRepository) error
	}
	tests := []struct {
		name     string
		args     args
		mockFunc func(sqlMock sqlmock.Sqlmock)
		expectedErr error
	}{
		{
			name: "success & commit",
			args: args{
				fn: func(tx repository.DemoRepository) error {
					return nil
				},
			},
			mockFunc: func(sqlMock sqlmock.Sqlmock) {
				sqlMock.ExpectBegin()
				sqlMock.ExpectCommit()
			},
		},
		{
			name: "failed & rollback",
			args: args{
				fn: func(tx repository.DemoRepository) error {
					return errors.New("error")
				},
			},
			mockFunc: func(sqlMock sqlmock.Sqlmock) {
				sqlMock.ExpectBegin()
				sqlMock.ExpectRollback()
			},
			expectedErr: errors.New("error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFunc != nil {
				tt.mockFunc(sqlMock)
			}
			err := dbc.Transaction(tt.args.fn)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func GetMockDBClient() (*DBClient, sqlmock.Sqlmock, error) {
	sqlDB, sqlMock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	db, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn:                      sqlDB,
			SkipInitializeWithVersion: true,
		},
	), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	return NewDBClient(db), sqlMock, nil
}
