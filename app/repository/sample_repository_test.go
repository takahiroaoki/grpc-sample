package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/go-env/config"
	"github.com/takahiroaoki/go-env/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestSampleRepository_SelectOneUserByUserId_Success(t *testing.T) {
	t.Parallel()

	db, _ := gorm.Open(
		mysql.Open(config.NewDataBaseConfig().GetDataSourceName()),
		&gorm.Config{},
	)

	sampleRepository := NewSampleRepository(db)

	userId := "1"
	expected := &entity.User{
		Model: gorm.Model{
			ID: uint(1),
		},
		Email: "user@example.com",
	}

	actual, err := sampleRepository.SelectOneUserByUserId(userId)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestSampleRepository_SelectOneUserByUserId_Error(t *testing.T) {
	t.Parallel()

	db, _ := gorm.Open(
		mysql.Open(config.NewDataBaseConfig().GetDataSourceName()),
		&gorm.Config{},
	)

	sampleRepository := NewSampleRepository(db)

	userId := "invalid id"
	var expected *entity.User

	actual, err := sampleRepository.SelectOneUserByUserId(userId)
	if assert.Error(t, err) {
		assert.Equal(t, "record not found", err.Error())
		assert.Equal(t, expected, actual)
	}
}
