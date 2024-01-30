package logger

import "github.com/sllt/log"

var DefaultLogger = &errLogger{}

type errLogger struct {
}

func (l *errLogger) Logf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}
