package main


import "github.com/j-keck/plog"

func main() {
    log := plog.NewBroadcastLogger(
        plog.NewDefaultConsoleLogger(),
        plog.NewDefaultConsoleLogger(),
        plog.NewDefaultConsoleLogger(),
    )

    log.Info("startup")
    log.Debug("change to debug level")
    log.SetLevel(plog.Debug)
    log.Debug("level changed")
    log.Infof("2 + 2 = %d", 2 + 2)
}
