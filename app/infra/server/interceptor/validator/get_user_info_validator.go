package validator

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/takahiroaoki/grpc-sample/app/infra/pb"
)

func (v *Validator) ValidateGetUserInfoRequest(ctx context.Context, req *pb.GetUserInfoRequest) error {
	rules := make([]*validation.FieldRules, 0)
	rules = append(rules, validation.Field(&req.Id, validation.Required, is.Digit))
	return validation.ValidateStructWithContext(ctx, req, rules...)
}
