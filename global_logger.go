package plog

import (
	"sync"
)

type globalLogger struct {
	broadcastLogger
}

var singleton *globalLogger
var once sync.Once

// GlobalLogger create a singleton global logger.
// You need to attach a 'Logger' to one singleton.
//
// Example:
//
//    plog.GlobalLogger().Add(plog.NewDefaultConsoleLogger())
//
//    log := plog.GlobalLogger()
//    log.Info("easy peasy!")
//
func GlobalLogger() *globalLogger {

	once.Do(func() {
		singleton = new(globalLogger)
		initLogger(singleton)
	})

	return singleton
}
