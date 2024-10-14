package domerr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_domErrImpl_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		aei      *domErrImpl
		expected string
	}{
		{
			name: "Success",
			aei: &domErrImpl{
				err: errors.New("err"),
			},
			expected: "err",
		},
		{
			name:     "Success(*domErrImpl is nil)",
			aei:      nil,
			expected: "",
		},
		{
			name:     "Success(err is nil)",
			aei:      &domErrImpl{},
			expected: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.aei.Error())
		})
	}
}

func Test_domErrImpl_Cause(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		aei      *domErrImpl
		expected ErrorCause
	}{
		{
			name: "Success",
			aei: &domErrImpl{
				err:   errors.New("err"),
				cause: CAUSE_INTERNAL,
			},
			expected: CAUSE_INTERNAL,
		},
		{
			name:     "Success(*domErrImpl is nil)",
			aei:      nil,
			expected: CAUSE_UNDEFINED,
		},
		{
			name:     "Success(err is nil)",
			aei:      &domErrImpl{},
			expected: CAUSE_UNDEFINED,
		},
		{
			name: "Error(cause is not defined)",
			aei: &domErrImpl{
				err: errors.New("err"),
			},
			expected: CAUSE_UNDEFINED,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.aei.Cause())
		})
	}
}

func Test_domErrImpl_LogLevel(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		aei      *domErrImpl
		expected LogLevel
	}{
		{
			name: "Success",
			aei: &domErrImpl{
				err:      errors.New("err"),
				logLevel: LOG_LEVEL_INFO,
			},
			expected: LOG_LEVEL_INFO,
		},
		{
			name:     "Success(*domErrImpl is nil)",
			aei:      nil,
			expected: LOG_LEVEL_UNDEFINED,
		},
		{
			name:     "Success(err is nil)",
			aei:      &domErrImpl{},
			expected: LOG_LEVEL_UNDEFINED,
		},
		{
			name: "Error(logLevel is not defined)",
			aei: &domErrImpl{
				err: errors.New("err"),
			},
			expected: LOG_LEVEL_UNDEFINED,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.aei.LogLevel())
		})
	}
}
