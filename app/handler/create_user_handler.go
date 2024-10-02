package handler

import (
	"context"
	"errors"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/handler/validator"
	"github.com/takahiroaoki/grpc-sample/app/repository"
	"github.com/takahiroaoki/grpc-sample/app/service"
)

type createUserHandlerImpl struct {
	dr  repository.DemoRepository
	cus service.CreateUserService
}

func (h *createUserHandlerImpl) Execute(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	if h == nil {
		return nil, errors.New("*createUserHandlerImpl is nil")
	}

	if err := h.validate(ctx, req); err != nil {
		return nil, err
	}

	var (
		u   *entity.User
		err error
	)
	err = h.dr.Transaction(func(dr repository.DemoRepository) error {
		u, err = h.cus.CreateUser(dr, entity.User{
			Email: req.email,
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return &CreateUserResponse{
		id: strconv.FormatUint(uint64(u.ID), 10),
	}, nil
}

func (h *createUserHandlerImpl) validate(ctx context.Context, req *CreateUserRequest) error {
	if h == nil {
		return errors.New("*createUserHandlerImpl is nil")
	}
	rules := make([]*validation.FieldRules, 0)
	rules = append(rules, validation.Field(&req.email, validation.Required, validation.RuneLength(1, 320), validation.Match(validator.MailRegexp())))

	return validation.ValidateStructWithContext(ctx, req, rules...)
}
