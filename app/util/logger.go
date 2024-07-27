package util

import "log"

func init() {
	log.SetFlags(log.Lmicroseconds)
}

func generalLog(category string, v string) {
	log.Printf("[%v] %v", category, v)
}

func InfoLog(v string) {
	generalLog("INFO", v)
}

func WarnLog(v string) {
	generalLog("WARN", v)
}

func ErrorLog(v string) {
	generalLog("ERROR", v)
}

func PerfLog(v string) {
	generalLog("PERF", v)
}

func FatalLog(v string) {
	log.Fatalf("[Fatal] %v", v)
}
