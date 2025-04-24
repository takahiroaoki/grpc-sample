package domerr

import (
	"errors"
	"fmt"
)

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
	AddDescription(description string) DomErr
}

type domErr struct {
	err      error
	cause    ErrorCause
	logLevel LogLevel
}

func (de *domErr) Error() string {
	if de == nil || de.err == nil {
		return ""
	}
	return de.err.Error()
}

func (de *domErr) Cause() ErrorCause {
	if de == nil || de.err == nil {
		return CAUSE_UNDEFINED
	}
	return de.cause
}

func (de *domErr) LogLevel() LogLevel {
	if de == nil || de.err == nil {
		return LOG_LEVEL_UNDEFINED
	}
	return de.logLevel
}

func (de *domErr) AddDescription(description string) DomErr {
	if de == nil {
		de = &domErr{}
	}
	de.err = fmt.Errorf("%s: %w", description, de)
	return de
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
