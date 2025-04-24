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

	var errMap = map[domerr.ErrorCause]error{
		domerr.CAUSE_UNDEFINED:        status.Error(codes.Unknown, "unknown error"),
		domerr.CAUSE_INVALID_ARGUMENT: status.Error(codes.InvalidArgument, "invalid argument"),
		domerr.CAUSE_NOT_FOUND:        status.Error(codes.NotFound, "data not found"),
		domerr.CAUSE_INTERNAL:         status.Error(codes.Internal, "internal error"),
	}

	switch err.Cause() {
	case domerr.CAUSE_UNDEFINED:
		util.ErrorLogWithContext(ctx, fmt.Sprintf("error with no cause: %s", err.Error()))
		return errMap[err.Cause()]
	default:
		return errMap[err.Cause()]
	}
}
