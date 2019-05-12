package main


import "github.com/j-keck/plog"
import "flag"

func main() {
    //
    // flags
    //
    logTs := flag.Bool("log-timestamps", false, "log messages with timestamps")
    logLocation := flag.Bool("log-location", false, "log messages with caller location")
    flag.Parse()

    //
    // initialize / configure the logger
    //
    log := plog.NewConsoleLogger(" | ")

    // timestamp only when '-log-timestamps' flag is given
    if *logTs {
        log.AddLogFormatter(plog.TimestampUnixDate)
    }

    // log level
    log.AddLogFormatter(plog.Level)

    // location only when '-log-location' flag is given
    if *logLocation {
        log.AddLogFormatter(plog.Location)
    }

    // log message
    log.AddLogFormatter(plog.Message)


    //
    // action
    //
    log.Info("startup")
    log.Debug("change to debug level")
    log.SetLevel(plog.Debug)
    log.Debug("level changed")
    log.Infof("2 + 2 = %d", 2 + 2)
}
