package database

import "github.com/takahiroaoki/grpc-sample/app/domain/entity"

// users TBL model
type user struct {
	id    uint   `gorm:"primaryKey"`
	email string `gorm:"type:varchar(255);not null"`
}

func convertUserEntity(e entity.User) *user {
	return &user{
		id:    e.ID,
		email: e.Email,
	}
}

func convertUserSchema(s user) *entity.User {
	return &entity.User{
		ID:    s.id,
		Email: s.email,
	}
}
