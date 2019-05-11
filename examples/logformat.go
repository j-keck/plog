package main


import "github.com/j-keck/plog"

func main() {
    log := plog.NewConsoleLogger(" - ",
        plog.LevelFmt("(%-5s)"),
        plog.TimestampFmt("2006-01-02T15:04:05Z07:00"),
        plog.MessageFmt("%-20s"),
        plog.LocationFmt("%s[%d]"),

    )
    log.SetLogPrefix("[").SetLogSuffix("]")

    log.Info("startup")
    log.Debug("change to debug level")
    log.SetLevel(plog.Debug)
    log.Debug("level changed")
    log.Infof("2 + 2 = %d", 2 + 2)
}
