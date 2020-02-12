/*
Package plog is a logging library.


Checkout https://github.com/j-keck/plog for more information.

------------------------------------------------------------------------------------------------------------------------------------------------------

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


------------------------------------------------------------------------------------------------------------------------------------------------------

Example how to configure the log output:

    package main

    import  "github.com/j-keck/plog"

    func main() {
        log := plog.NewConsoleLogger(
            " - ",
            plog.Level,
            plog.TimestampMillis,
            plog.Message,
        )

        // set log prefix and suffix
        log.SetLogPrefix("[").SetLogSuffix("]")

        log.Info("startup")
        log.Debug("change to debug level")
        log.SetLevel(plog.Debug)
        log.Debug("level changed")
        log.Infof("2 + 2 = %d", 2 + 2)
    }

Output:

   [ INFO - May 11 19:01:21.574 - startup]
   [DEBUG - May 11 19:01:21.575 - level changed]
   [ INFO - May 11 19:01:21.575 - 2 + 2 = 4]

*/
package plog

var dropUnhandledMessages bool = false

// Disables the logging of unhandled messages.
//
// If you use a `BroadcastLogger` or a `GlobalLogger`
// without a attached logger, the log-messages were
// unhandled and plog will log a warning about this.
//
// Call `plog.DropUnhandledMessages()` to disable
// the warnings.
func DropUnhandledMessages() {
	dropUnhandledMessages = true
}


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
