package database

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
	"github.com/takahiroaoki/grpc-sample/app/testutil"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_DBClient_Transaction(t *testing.T) {
	t.Parallel()
	type args struct {
		fn func(tx repository.DemoRepository) error
	}
	tests := []struct {
		name        string
		args        args
		mockFunc    func(sqlMock sqlmock.Sqlmock)
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
			t.Parallel()
			dbc, sqlMock, err := GetMockDBClient()
			defer func() {
				_ = dbc.CloseDB()
			}()
			assert.NoError(t, err)
			if tt.mockFunc != nil {
				tt.mockFunc(sqlMock)
			}
			tranErr := dbc.Transaction(tt.args.fn)
			assert.Equal(t, tt.expectedErr, tranErr)
			assert.Nil(t, sqlMock.ExpectationsWereMet())
		})
	}
}

func Test_DBClient_SelectOneUserByUserId(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		userId string
	}
	tests := []struct {
		name        string
		args        args
		mockFunc    func(sqlMock sqlmock.Sqlmock)
		assertions  assert.ErrorAssertionFunc
		expectedUsr *entity.User
		expectedErr domerr.DomErr
	}{
		{
			name: "success",
			args: args{
				ctx:    context.Background(),
				userId: "1",
			},
			mockFunc: func(sqlMock sqlmock.Sqlmock) {
				sqlMock.ExpectQuery("SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT ?").
					WithArgs("1", 1).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "email"}).AddRow(uint(1), "sample@example.com"),
					)
			},
			expectedUsr: &entity.User{
				ID:    1,
				Email: "sample@example.com",
			},
			assertions: assert.NoError,
		},
		{
			name: "error > record not found",
			args: args{
				ctx:    context.Background(),
				userId: "1",
			},
			mockFunc: func(sqlMock sqlmock.Sqlmock) {
				sqlMock.ExpectQuery("SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT ?").
					WithArgs("1", 1).
					WillReturnError(gorm.ErrRecordNotFound)
			},
			expectedUsr: nil,
			expectedErr: domerr.NewDomErr(errors.New("DBClient.SelectOneUserByUserId: record not found"), domerr.CAUSE_NOT_FOUND, domerr.LOG_LEVEL_INFO),
			assertions:  assert.Error,
		},
		{
			name: "error > other errors",
			args: args{
				ctx:    context.Background(),
				userId: "1",
			},
			mockFunc: func(sqlMock sqlmock.Sqlmock) {
				sqlMock.ExpectQuery("SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT ?").
					WithArgs("1", 1).
					WillReturnError(gorm.ErrInvalidDB)
			},
			expectedUsr: nil,
			expectedErr: domerr.NewDomErr(errors.New("DBClient.SelectOneUserByUserId: invalid db"), domerr.CAUSE_INTERNAL, domerr.LOG_LEVEL_ERROR),
			assertions:  assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			dbc, sqlMock, err := GetMockDBClient()
			defer func() {
				_ = dbc.CloseDB()
			}()
			assert.NoError(t, err)
			if tt.mockFunc != nil {
				tt.mockFunc(sqlMock)
			}
			gotUsr, gotErr := dbc.SelectOneUserByUserId(tt.args.ctx, tt.args.userId)
			assert.Equal(t, tt.expectedUsr, gotUsr)
			tt.assertions(t, gotErr)
			if gotErr != nil {
				assert.True(t, testutil.SameDomainErrors(tt.expectedErr, gotErr))
			}
			assert.Nil(t, sqlMock.ExpectationsWereMet())
		})
	}
}

func GetMockDBClient() (*DBClient, sqlmock.Sqlmock, error) {
	sqlDB, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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
