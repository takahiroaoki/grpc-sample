package service

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
)

type createUserService struct {
	cur createUserRepository
}

type createUserRepository interface {
	// Transaction(fn func(tx createUserRepository) error) error
	CreateOneUser(ctx context.Context, u entity.User) (*entity.User, domerr.DomErr)
}

func (s *createUserService) CreateUser(ctx context.Context, u entity.User) (*entity.User, domerr.DomErr) {
	return s.cur.CreateOneUser(ctx, u)
}

func NewCreateUserService(cur createUserRepository) *createUserService {
	return &createUserService{
		cur: cur,
	}
}
