package main

import "github.com/j-keck/plog"

func main() {
    log := plog.NewBroadcastLogger(
        plog.NewConsoleLogger(" | ", plog.Message, plog.Level),
        plog.NewConsoleLogger(" - ", plog.Message, plog.TimestampUnixDate),
        plog.NewConsoleLogger(" / ", plog.Message, plog.TimestampMillis),
    )

    log.Info("startup")
    log.Debug("change to debug level")
    log.SetLevel(plog.Debug)
    log.Debug("level changed")
}
