package logger

import (
	"log"
)

func Init(){
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Printf("[LOGGER] initialized")
}

func Log(format string, v ... any){
	if len(v) == 0 {
		log.Printf("[LOG] %s", format)
		return
	}
	log.Printf("[LOG] "+format, v...)
}