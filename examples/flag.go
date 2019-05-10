package main

import "github.com/j-keck/plog"
import "flag"


func main() {
    log := plog.NewDefaultConsoleLogger()

    logLevel := plog.Info
    plog.FlagDebugVar(&logLevel,  "v", "debug")
    plog.FlagTraceVar(&logLevel, "vv", "trace")
    flag.Parse()

    log.SetLevel(logLevel)

    log.Info("info")
    log.Debug("debug")
    log.Trace("trace")
}
