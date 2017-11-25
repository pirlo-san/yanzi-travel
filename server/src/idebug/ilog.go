package idebug

import (
	"log"
)

var debugOn bool = false

func SetDebugMode(on bool) {
	debugOn = on
	log.Printf("debug mode %t\n", on)
}

func InfoLog(format string, arg ...interface{}) {
	if debugOn == false {
		return
	}

	log.Printf(format, arg...)
}

func ErrorLog(format string, arg ...interface{}) {

	log.Printf(format, arg...)
}
