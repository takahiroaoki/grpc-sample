package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
)

func Test_convertUserEntity(t *testing.T) {
	e := entity.User{
		ID:    1,
		Email: "email",
	}
	expected := user{
		ID:    1,
		Email: "email",
	}
	assert.Equal(t, expected, convertUserEntity(e))
}

func Test_convertUserSchema(t *testing.T) {
	s := user{
		ID:    1,
		Email: "email",
	}
	expected := entity.User{
		ID:    1,
		Email: "email",
	}
	assert.Equal(t, expected, convertUserSchema(s))
}
