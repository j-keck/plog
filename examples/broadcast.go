package main

import "github.com/j-keck/plog"

func main() {
    log := plog.NewBroadcastLogger(
        plog.NewConsoleLogger(),
        plog.NewConsoleLogger(),
        plog.NewConsoleLogger(),
    )

    log.Info("startup")
    log.Debug("change to debug level")
    log.SetLevel(plog.Debug)
    log.Debug("level changed")
}
