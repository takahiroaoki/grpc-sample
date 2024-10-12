package util

import "errors"

type ErrorCause int

const (
	CAUSE_UNDEFINED ErrorCause = iota
	CAUSE_INVALID_ARGUMENT
	CAUSE_NOT_FOUND
	CAUSE_INTERNAL
)

type AppError interface {
	Error() string
	Cause() ErrorCause
	LogLevel() LogLevel
	Equals(err AppError) bool
}

type appErrorImpl struct {
	err      error
	cause    ErrorCause
	logLevel LogLevel
}

func (aei *appErrorImpl) Error() string {
	if aei == nil || aei.err == nil {
		return ""
	}
	return aei.err.Error()
}

func (aei *appErrorImpl) Cause() ErrorCause {
	if aei == nil || aei.err == nil {
		return CAUSE_UNDEFINED
	}
	return aei.cause
}

func (aei *appErrorImpl) LogLevel() LogLevel {
	if aei == nil || aei.err == nil {
		return LOG_LEVEL_UNDEFINED
	}
	return aei.logLevel
}

func (aei *appErrorImpl) Equals(err AppError) bool {
	return (aei.err.Error() == err.Error()) && (aei.cause == err.Cause()) && (aei.logLevel == err.LogLevel())
}

func NewAppError(err error, cause ErrorCause, logLevel LogLevel) AppError {
	if err == nil {
		return nil
	}
	return &appErrorImpl{
		err:      err,
		cause:    cause,
		logLevel: logLevel,
	}
}

func NewAppErrorFromMsg(msg string, cause ErrorCause, logLevel LogLevel) AppError {
	return NewAppError(errors.New(msg), cause, logLevel)
}
