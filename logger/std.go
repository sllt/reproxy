package logger

import (
	L "github.com/sllt/log"
	"log"
)

var (
	stdLogger *log.Logger
)

func StdLogger() *log.Logger {

	if stdLogger == nil {
		stdLogger = L.Default().StandardLog(L.StandardLogOptions{
			ForceLevel: L.WarnLevel,
		})
	}

	return stdLogger
}
