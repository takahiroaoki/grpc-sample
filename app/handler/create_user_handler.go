package handler

import (
	"context"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/takahiroaoki/grpc-sample/app/entity"
	"github.com/takahiroaoki/grpc-sample/app/handler/validator"
	"github.com/takahiroaoki/grpc-sample/app/repository"
	"github.com/takahiroaoki/grpc-sample/app/service"
	"github.com/takahiroaoki/grpc-sample/app/util"
)

type createUserHandlerImpl struct {
	dr  repository.DemoRepository
	cus service.CreateUserService
}

func (h *createUserHandlerImpl) process(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, util.AppError) {
	if h == nil {
		return nil, util.NewAppErrorFromMsg("*createUserHandlerImpl is nil", util.CAUSE_INTERNAL, util.LOG_LEVEL_ERROR)
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
		appErr, ok := err.(util.AppError)
		if !ok {
			return nil, util.NewAppError(err, util.CAUSE_INTERNAL, util.LOG_LEVEL_ERROR)
		}
		return nil, appErr
	}
	return &CreateUserResponse{
		id: strconv.FormatUint(uint64(u.ID), 10),
	}, nil
}

func (h *createUserHandlerImpl) validate(ctx context.Context, req *CreateUserRequest) util.AppError {
	if h == nil {
		return util.NewAppErrorFromMsg("*createUserHandlerImpl is nil", util.CAUSE_INTERNAL, util.LOG_LEVEL_ERROR)
	}
	rules := make([]*validation.FieldRules, 0)
	rules = append(rules, validation.Field(&req.email, validation.Required, validation.RuneLength(1, 320), validation.Match(validator.MailRegexp())))

	return util.NewAppError(
		validation.ValidateStructWithContext(ctx, req, rules...),
		util.CAUSE_INVALID_ARGUMENT,
		util.LOG_LEVEL_INFO,
	)
}

func NewCreateUserHandler(dr repository.DemoRepository, cus service.CreateUserService) CreateUserHandler {
	return &createUserHandlerImpl{
		dr:  dr,
		cus: cus,
	}
}
