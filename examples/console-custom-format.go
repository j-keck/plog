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
