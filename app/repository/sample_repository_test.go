package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/go-env/app/entity"
	"github.com/takahiroaoki/go-env/app/testutil"
)

func TestSampleRepository_SelectOneUserByUserId_Success(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	sampleRepository := NewSampleRepository()

	userId := "1"
	expected := "user@example.com"

	sampleRepository.CreateOneUser(db, entity.User{
		Email: "user@example.com",
	})

	actual, err := sampleRepository.SelectOneUserByUserId(db, userId)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual.Email)
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

func TestSampleRepository_CreateOneUser_Success(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	sampleRepository := NewSampleRepository()

	u := entity.User{
		Email: "user@example.com",
	}
	expected := "user@example.com"

	actual, err := sampleRepository.CreateOneUser(db, u)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual.Email)
	}
}

func TestSampleRepository_CreateOneUser_Error(t *testing.T) {
	t.Parallel()

	db, _ := testutil.GetDatabase()

	sampleRepository := NewSampleRepository()

	u := entity.User{
		Email: "123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890@example.com",
	}
	var expected *entity.User

	actual, err := sampleRepository.CreateOneUser(db, u)

	if assert.Error(t, err) {
		assert.Equal(t, "Error 1406 (22001): Data too long for column 'email' at row 1", err.Error())
		assert.Equal(t, expected, actual)
	}

}
