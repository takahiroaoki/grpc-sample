package handler

import (
	"context"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/takahiroaoki/go-env/app/constant"
	"github.com/takahiroaoki/go-env/app/entity"
	"github.com/takahiroaoki/go-env/app/pb"
	"github.com/takahiroaoki/go-env/app/service"
	"gorm.io/gorm"
)

type createUserHandler struct {
	db                *gorm.DB
	createUserService service.CreateUserService
}

func (h *createUserHandler) execute(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if err := h.validate(ctx, req); err != nil {
		return nil, err
	}

	var (
		u   *entity.User
		err error
	)
	err = h.db.Transaction(func(tx *gorm.DB) error {
		u, err = h.createUserService.CreateUser(tx, entity.User{
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

func (h *createUserHandler) validate(ctx context.Context, req *pb.CreateUserRequest) error {
	rules := make([]*validation.FieldRules, 0)
	rules = append(rules, validation.Field(&req.Email, validation.Required, validation.RuneLength(1, 320), validation.Match(constant.MailRegexp())))

	return validation.ValidateStructWithContext(ctx, req, rules...)
}

func NewCreateUserHandler(db *gorm.DB, createUserService service.CreateUserService) Handler[*pb.CreateUserRequest, *pb.CreateUserResponse] {
	return &createUserHandler{
		db:                db,
		createUserService: createUserService,
	}
}
