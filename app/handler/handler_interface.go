package handler

import "context"

type Handler[Req, Res any] interface {
	execute(ctx context.Context, req Req) (Res, error)
	validate(ctx context.Context, req Req) error
}
