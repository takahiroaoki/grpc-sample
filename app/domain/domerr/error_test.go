package domerr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_domErr_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		de       *domErr
		expected string
	}{
		{
			name: "Success",
			de: &domErr{
				err: errors.New("err"),
			},
			expected: "err",
		},
		{
			name:     "Success(*domErr is nil)",
			de:       nil,
			expected: "",
		},
		{
			name:     "Success(err is nil)",
			de:       &domErr{},
			expected: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, tt.de.Error())
		})
	}
}

func Test_domErr_Cause(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		de       *domErr
		expected ErrorCause
	}{
		{
			name: "Success",
			de: &domErr{
				err:   errors.New("err"),
				cause: CAUSE_INTERNAL,
			},
			expected: CAUSE_INTERNAL,
		},
		{
			name:     "Success(*domErr is nil)",
			de:       nil,
			expected: CAUSE_UNDEFINED,
		},
		{
			name:     "Success(err is nil)",
			de:       &domErr{},
			expected: CAUSE_UNDEFINED,
		},
		{
			name: "Error(cause is not defined)",
			de: &domErr{
				err: errors.New("err"),
			},
			expected: CAUSE_UNDEFINED,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, tt.de.Cause())
		})
	}
}

func Test_domErr_LogLevel(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		de       *domErr
		expected LogLevel
	}{
		{
			name: "Success",
			de: &domErr{
				err:      errors.New("err"),
				logLevel: LOG_LEVEL_INFO,
			},
			expected: LOG_LEVEL_INFO,
		},
		{
			name:     "Success(*domErr is nil)",
			de:       nil,
			expected: LOG_LEVEL_UNDEFINED,
		},
		{
			name:     "Success(err is nil)",
			de:       &domErr{},
			expected: LOG_LEVEL_UNDEFINED,
		},
		{
			name: "Error(logLevel is not defined)",
			de: &domErr{
				err: errors.New("err"),
			},
			expected: LOG_LEVEL_UNDEFINED,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, tt.de.LogLevel())
		})
	}
}

func Test_domErr_AddContext(t *testing.T) {
	t.Parallel()

	type args struct {
		description string
	}

	tests := []struct {
		name     string
		de       *domErr
		args     args
		expected *domErr
	}{
		{
			name: "Success",
			de: &domErr{
				err:      errors.New("err"),
				logLevel: LOG_LEVEL_INFO,
			},
			args: args{
				description: "description",
			},
			expected: &domErr{
				err:      errors.New("description: err"),
				logLevel: LOG_LEVEL_INFO,
			},
		},
		{
			name: "Success(*domErr is nil)",
			de:   nil,
			args: args{
				description: "description",
			},
			expected: &domErr{
				err: errors.New("description: "),
			},
		},
		{
			name: "Success(err is nil)",
			de:   &domErr{},
			args: args{
				description: "description",
			},
			expected: &domErr{
				err: errors.New("description: "),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.de.AddDescription(tt.args.description)
			assert.EqualError(t, tt.expected, got.Error())
		})
	}
}
