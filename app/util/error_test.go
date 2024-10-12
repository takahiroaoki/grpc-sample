package util

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_appErrorImpl_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		aei      *appErrorImpl
		expected string
	}{
		{
			name: "Success",
			aei: &appErrorImpl{
				err: errors.New("err"),
			},
			expected: "err",
		},
		{
			name:     "Success(*appErrorImpl is nil)",
			aei:      nil,
			expected: "",
		},
		{
			name:     "Success(err is nil)",
			aei:      &appErrorImpl{},
			expected: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.aei.Error())
		})
	}
}

func Test_appErrorImpl_Cause(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		aei      *appErrorImpl
		expected ErrorCause
	}{
		{
			name: "Success",
			aei: &appErrorImpl{
				err:   errors.New("err"),
				cause: CAUSE_INTERNAL,
			},
			expected: CAUSE_INTERNAL,
		},
		{
			name:     "Success(*appErrorImpl is nil)",
			aei:      nil,
			expected: CAUSE_UNDEFINED,
		},
		{
			name:     "Success(err is nil)",
			aei:      &appErrorImpl{},
			expected: CAUSE_UNDEFINED,
		},
		{
			name: "Error(cause is not defined)",
			aei: &appErrorImpl{
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

func Test_appErrorImpl_LogLevel(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		aei      *appErrorImpl
		expected LogLevel
	}{
		{
			name: "Success",
			aei: &appErrorImpl{
				err:      errors.New("err"),
				logLevel: LOG_LEVEL_INFO,
			},
			expected: LOG_LEVEL_INFO,
		},
		{
			name:     "Success(*appErrorImpl is nil)",
			aei:      nil,
			expected: LOG_LEVEL_UNDEFINED,
		},
		{
			name:     "Success(err is nil)",
			aei:      &appErrorImpl{},
			expected: LOG_LEVEL_UNDEFINED,
		},
		{
			name: "Error(logLevel is not defined)",
			aei: &appErrorImpl{
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

func Test_appErrorImpl_Equals(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		aei1      *appErrorImpl
		aei2      *appErrorImpl
		assertion assert.BoolAssertionFunc
	}{
		{
			name: "True",
			aei1: &appErrorImpl{
				err:      errors.New("err"),
				cause:    CAUSE_NOT_FOUND,
				logLevel: LOG_LEVEL_INFO,
			},
			aei2: &appErrorImpl{
				err:      errors.New("err"),
				cause:    CAUSE_NOT_FOUND,
				logLevel: LOG_LEVEL_INFO,
			},
			assertion: assert.True,
		},
		{
			name: "False(error message)",
			aei1: &appErrorImpl{
				err:      errors.New("err1"),
				cause:    CAUSE_NOT_FOUND,
				logLevel: LOG_LEVEL_INFO,
			},
			aei2: &appErrorImpl{
				err:      errors.New("err2"),
				cause:    CAUSE_NOT_FOUND,
				logLevel: LOG_LEVEL_INFO,
			},
			assertion: assert.False,
		},
		{
			name: "False(cause)",
			aei1: &appErrorImpl{
				err:      errors.New("err"),
				cause:    CAUSE_NOT_FOUND,
				logLevel: LOG_LEVEL_INFO,
			},
			aei2: &appErrorImpl{
				err:      errors.New("err"),
				cause:    CAUSE_INTERNAL,
				logLevel: LOG_LEVEL_INFO,
			},
			assertion: assert.False,
		},
		{
			name: "False(logLevel)",
			aei1: &appErrorImpl{
				err:      errors.New("err"),
				cause:    CAUSE_NOT_FOUND,
				logLevel: LOG_LEVEL_INFO,
			},
			aei2: &appErrorImpl{
				err:      errors.New("err"),
				cause:    CAUSE_NOT_FOUND,
				logLevel: LOG_LEVEL_WARN,
			},
			assertion: assert.False,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, tt.aei1.Equals(tt.aei2))
		})
	}
}
