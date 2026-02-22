package logger

import (
	"log"
)

func Init(lvl string){
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Printf("[LOGGER] initialized")
}