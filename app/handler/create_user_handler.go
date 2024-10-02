package handler

import (
	"context"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/handler/validator"
	"github.com/takahiroaoki/grpc-sample/app/pb"
	"github.com/takahiroaoki/grpc-sample/app/repository"
	"github.com/takahiroaoki/grpc-sample/app/service"
)

type createUserHandlerImpl struct {
	dr  repository.DemoRepository
	cus service.CreateUserService
}

func (h *createUserHandlerImpl) execute(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if err := h.validate(ctx, req); err != nil {
		return nil, err
	}

	var (
		u   *entity.User
		err error
	)
	err = h.dr.Transaction(func(dr repository.DemoRepository) error {
		u, err = h.cus.CreateUser(dr, entity.User{
			Email: req.GetEmail(),
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		Id: strconv.FormatUint(uint64(u.ID), 10),
	}, nil
}

func (h *createUserHandlerImpl) validate(ctx context.Context, req *pb.CreateUserRequest) error {
	rules := make([]*validation.FieldRules, 0)
	rules = append(rules, validation.Field(&req.Email, validation.Required, validation.RuneLength(1, 320), validation.Match(validator.MailRegexp())))

	return validation.ValidateStructWithContext(ctx, req, rules...)
}
