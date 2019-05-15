package log

import (
	"log"
)

func Info(format string, v ...interface{}) {
	format2 := "[INFO] " + format 
	log.Printf(format2, v...)
}

//func Debug(message string) {
//	log.Print("[DEBUG] " + message)
//}
func Debug(format string, v ...interface{}) {
	format2 := "[DEBUG] " + format 
	if (len(v) > 0 ) {
		log.Printf(format2, v...)
	} else {
		log.Print(format2)
	}
}