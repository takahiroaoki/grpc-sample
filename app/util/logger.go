package util

import "log"

func InfoLog(v string) {
	log.Printf("[INFO] %v", v)
}

func WarnLog(v string) {
	log.Printf("[Warn] %v", v)
}

func ErrorLog(v string) {
	log.Printf("[ERROR] %v", v)
}

func FatalLog(v string) {
	log.Fatalf("[Fatal] %v", v)
}
