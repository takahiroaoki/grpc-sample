package database

import "github.com/takahiroaoki/grpc-sample/app/domain/entity"

// users TBL model
type user struct {
	ID    uint   `gorm:"primaryKey"`
	Email string `gorm:"type:varchar(255);not null"`
}

func convertUserEntity(e entity.User) *user {
	return &user{
		ID:    e.ID,
		Email: e.Email,
	}
}

func convertUserSchema(s user) *entity.User {
	return &entity.User{
		ID:    s.ID,
		Email: s.Email,
	}
}
