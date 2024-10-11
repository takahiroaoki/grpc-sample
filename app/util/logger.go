package util

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/takahiroaoki/grpc-sample/app/constant"
)

type LogLevel int

const (
	LOG_LEVEL_UNDEFINED LogLevel = iota
	LOG_LEVEL_INFO
	LOG_LEVEL_WARN
	LOG_LEVEL_ERROR
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.LUTC|log.Lmicroseconds)
}

func generalLog(category string, v string) {
	logger.Printf("[%v] %v", category, v)
}

func addContextInfo(ctx context.Context, v string) string {
	for _, key := range constant.ContextKeysForLog() {
		val := ctx.Value(key)
		if val != nil {
			v = fmt.Sprintf("%v: %v, %v", key, val, v)
		}
	}
	return v
}

func InfoLog(v string) {
	generalLog("INFO", v)
}

func InfoLogWithContext(ctx context.Context, v string) {
	InfoLog(addContextInfo(ctx, v))
}

func WarnLog(v string) {
	generalLog("WARN", v)
}

func WarnLogWithContext(ctx context.Context, v string) {
	WarnLog(addContextInfo(ctx, v))
}

func ErrorLog(v string) {
	generalLog("ERROR", v)
}

func ErrorLogWithContext(ctx context.Context, v string) {
	ErrorLog(addContextInfo(ctx, v))
}

func PerfLog(v string) {
	generalLog("PERF", v)
}

func PerfLogWithContext(ctx context.Context, v string) {
	PerfLog(addContextInfo(ctx, v))
}

func FatalLog(v string) {
	logger.Fatalf("[Fatal] %v", v)
}
