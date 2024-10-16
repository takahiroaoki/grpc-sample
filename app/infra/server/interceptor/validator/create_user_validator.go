package validator

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/takahiroaoki/grpc-sample/app/infra/pb"
)

func (v *Validator) ValidateCreateUserRequest(ctx context.Context, req *pb.CreateUserRequest) error {
	rules := make([]*validation.FieldRules, 0)
	rules = append(rules, validation.Field(&req.Email, validation.Required, validation.RuneLength(1, 320), validation.Match(mailRegexp)))
	return validation.ValidateStructWithContext(ctx, req, rules...)
}
