package main

import "github.com/j-keck/plog"


func main() {
    log := plog.NewConsoleLogger()
    log.SetLogFields(
        plog.Level("%-5s"),
        plog.Timestamp("Mon Jan 2 15:04:05"),
        plog.Message("%s"),
    ).SetLogSeparator(" - ")

    log.Info("startup")
    log.Debug("change to debug level")
    log.SetLevel(plog.Debug)
    log.Debug("level changed")
}
