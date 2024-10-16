package server

import (
	"context"
	"fmt"

	"github.com/takahiroaoki/grpc-sample/app/domain/domerr"
	"github.com/takahiroaoki/grpc-sample/app/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func handleError(ctx context.Context, err domerr.DomErr) error {
	if err == nil {
		return nil
	}

	switch err.LogLevel() {
	case domerr.LOG_LEVEL_INFO:
		util.InfoLogWithContext(ctx, err.Error())
	case domerr.LOG_LEVEL_WARN:
		util.WarnLogWithContext(ctx, err.Error())
	case domerr.LOG_LEVEL_ERROR:
		util.ErrorLogWithContext(ctx, err.Error())
	default:
		util.ErrorLogWithContext(ctx, fmt.Sprintf("error with no log level: %s", err.Error()))
	}

	switch err.Cause() {
	case domerr.CAUSE_INVALID_ARGUMENT:
		return status.Error(codes.InvalidArgument, err.Error())
	case domerr.CAUSE_NOT_FOUND:
		return status.Error(codes.NotFound, err.Error())
	case domerr.CAUSE_INTERNAL:
		return status.Error(codes.Internal, err.Error())
	default:
		util.ErrorLogWithContext(ctx, fmt.Sprintf("error with no cause: %s", err.Error()))
		return status.Error(codes.Internal, err.Error())
	}
}
