package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/go-env/entity"
	"github.com/takahiroaoki/go-env/testutil"
	"gorm.io/gorm"
)

func TestSampleRepository_SelectOneUserByUserId_Success(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	sampleRepository := NewSampleRepository()

	userId := "1"
	expected := &entity.User{
		Model: gorm.Model{
			ID: uint(1),
		},
		Email: "user@example.com",
	}

	actual, err := sampleRepository.SelectOneUserByUserId(db, userId)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestSampleRepository_SelectOneUserByUserId_Error(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	sampleRepository := NewSampleRepository()

	userId := "invalid id"
	var expected *entity.User

	actual, err := sampleRepository.SelectOneUserByUserId(db, userId)
	if assert.Error(t, err) {
		assert.Equal(t, "record not found", err.Error())
		assert.Equal(t, expected, actual)
	}
}
