/*
Package plog is a logging library.

You can log over a console or over a go chan.

--------------------------------------------------------------------------------
Example how to log over a console:

    package main

    import  "github.com/j-keck/plog"

    func main() {
        log := plog.NewDefaultConsoleLogger()

        log.Info("startup")
        log.Debug("change to debug level")
        log.SetLevel(plog.Debug)
        log.Debug("level changed")
        log.Infof("2 + 2 = %d", 2 + 2)
    }


Output:

    Sat May 11 18:18:08 CEST 2019 |  INFO | startup
    Sat May 11 18:18:08 CEST 2019 | DEBUG | level changed
    Sat May 11 18:18:08 CEST 2019 |  INFO | 2 + 2 = 4

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
