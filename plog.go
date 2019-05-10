/*
Package plog is a simple logging library.


You can log over a console or over a go chan.


Example how to log over a console:

    package main

    import "github.com/j-keck/plog"

    func main() {
      log := plog.NewConsoleLogger()
      log.Infof("2 + 2 = %d", 2 + 2)
    }
*/
package plog

// Every `Logger' implements this interface.
//
// This are the basic functions every logger implements.
// Logger implementations can have more functions.
type Logger interface {
	SetLevel(LogLevel)
	Trace(...interface{})
	Tracef(string, ...interface{})
	Debug(...interface{})
	Debugf(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Warn(...interface{})
	Warnf(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
}
