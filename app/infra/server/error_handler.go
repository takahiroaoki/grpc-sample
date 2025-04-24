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

	var errMap = map[domerr.ErrorCause]string{
		domerr.CAUSE_UNDEFINED:        "internal error",
		domerr.CAUSE_INVALID_ARGUMENT: "invalid argument",
		domerr.CAUSE_NOT_FOUND:        "data not found",
		domerr.CAUSE_INTERNAL:         "internal error",
	}

	switch err.Cause() {
	case domerr.CAUSE_INVALID_ARGUMENT:
		return status.Error(codes.InvalidArgument, errMap[err.Cause()])
	case domerr.CAUSE_NOT_FOUND:
		return status.Error(codes.NotFound, errMap[err.Cause()])
	case domerr.CAUSE_INTERNAL:
		return status.Error(codes.Internal, errMap[err.Cause()])
	default:
		util.ErrorLogWithContext(ctx, fmt.Sprintf("error with no cause: %s", err.Error()))
		return status.Error(codes.Internal, errMap[err.Cause()])
	}
}
