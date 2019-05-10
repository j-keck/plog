package main

import  "github.com/j-keck/plog"

func main() {
    log := plog.NewDefaultConsoleLogger()

    log.Info("startup")
    log.Debug("change to debug level")
    log.SetLevel(plog.Debug)
    log.Debug("level changed")
}
