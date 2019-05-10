package main

import "github.com/j-keck/plog"

func main() {
    log := plog.NewConsoleLogger(" - ",
        plog.TimestampFmt("2006-01-02T15:04:05Z07:00"),
        plog.LevelFmt("(%-5s)"),
        plog.Location,
        plog.Message,
    )

    log.Info("startup")
    log.Debug("change to debug level")
    log.SetLevel(plog.Debug)
    log.Debug("level changed")
}
