package domerr

import "errors"

type ErrorCause int

const (
	CAUSE_UNDEFINED ErrorCause = iota
	CAUSE_INVALID_ARGUMENT
	CAUSE_NOT_FOUND
	CAUSE_INTERNAL
)

type LogLevel int

const (
	LOG_LEVEL_UNDEFINED LogLevel = iota
	LOG_LEVEL_INFO
	LOG_LEVEL_WARN
	LOG_LEVEL_ERROR
)

type DomErr interface {
	Error() string
	Cause() ErrorCause
	LogLevel() LogLevel
}

type domErr struct {
	err      error
	cause    ErrorCause
	logLevel LogLevel
}

func (aei *domErr) Error() string {
	if aei == nil || aei.err == nil {
		return ""
	}
	return aei.err.Error()
}

func (aei *domErr) Cause() ErrorCause {
	if aei == nil || aei.err == nil {
		return CAUSE_UNDEFINED
	}
	return aei.cause
}

func (aei *domErr) LogLevel() LogLevel {
	if aei == nil || aei.err == nil {
		return LOG_LEVEL_UNDEFINED
	}
	return aei.logLevel
}

func NewDomErr(err error, cause ErrorCause, logLevel LogLevel) DomErr {
	if err == nil {
		return nil
	}
	return &domErr{
		err:      err,
		cause:    cause,
		logLevel: logLevel,
	}
}

func NewDomErrFromMsg(msg string, cause ErrorCause, logLevel LogLevel) DomErr {
	return NewDomErr(errors.New(msg), cause, logLevel)
}
