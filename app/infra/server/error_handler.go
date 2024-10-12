package server

import (
	"context"

	"github.com/takahiroaoki/grpc-sample/app/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func handleError(ctx context.Context, err util.AppError) error {
	if err == nil {
		return nil
	}

	switch err.LogLevel() {
	case util.LOG_LEVEL_INFO:
		util.InfoLogWithContext(ctx, err.Error())
	case util.LOG_LEVEL_WARN:
		util.WarnLogWithContext(ctx, err.Error())
	case util.LOG_LEVEL_ERROR:
		util.ErrorLogWithContext(ctx, err.Error())
	default:
	}

	switch err.Cause() {
	case util.CAUSE_INVALID_ARGUMENT:
		return status.Error(codes.InvalidArgument, err.Error())
	case util.CAUSE_NOT_FOUND:
		return status.Error(codes.NotFound, err.Error())
	case util.CAUSE_INTERNAL:
		return status.Error(codes.Internal, err.Error())
	default:
		return nil
	}
}
