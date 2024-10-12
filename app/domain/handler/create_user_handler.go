package handler

import (
	"context"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/domain/entity"
	"github.com/takahiroaoki/grpc-sample/app/domain/handler/validator"
	"github.com/takahiroaoki/grpc-sample/app/domain/repository"
	"github.com/takahiroaoki/grpc-sample/app/domain/service"
)

type createUserHandlerImpl struct {
	dr  repository.DemoRepository
	cus service.CreateUserService
}

func (h *createUserHandlerImpl) process(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, domerr.DomErr) {
	if h == nil {
		return nil, domerr.NewDomErrFromMsg("*createUserHandlerImpl is nil", domerr.CAUSE_INTERNAL, domerr.LOG_LEVEL_ERROR)
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
		appErr, ok := err.(domerr.DomErr)
		if !ok {
			return nil, domerr.NewDomErr(err, domerr.CAUSE_INTERNAL, domerr.LOG_LEVEL_ERROR)
		}
		return nil, appErr
	}
	return &CreateUserResponse{
		id: strconv.FormatUint(uint64(u.ID), 10),
	}, nil
}

func (h *createUserHandlerImpl) validate(ctx context.Context, req *CreateUserRequest) domerr.DomErr {
	if h == nil {
		return domerr.NewDomErrFromMsg("*createUserHandlerImpl is nil", domerr.CAUSE_INTERNAL, domerr.LOG_LEVEL_ERROR)
	}
	rules := make([]*validation.FieldRules, 0)
	rules = append(rules, validation.Field(&req.email, validation.Required, validation.RuneLength(1, 320), validation.Match(validator.MailRegexp())))

	return domerr.NewDomErr(
		validation.ValidateStructWithContext(ctx, req, rules...),
		domerr.CAUSE_INVALID_ARGUMENT,
		domerr.LOG_LEVEL_INFO,
	)
}

func NewCreateUserHandler(dr repository.DemoRepository, cus service.CreateUserService) CreateUserHandler {
	return &createUserHandlerImpl{
		dr:  dr,
		cus: cus,
	}
}
