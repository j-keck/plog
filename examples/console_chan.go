package main

import "time"
import "fmt"
import "github.com/j-keck/plog"

func main() {
    log := plog.NewConsoleLogger()
    defer log.WaitForSubscribers(1 * time.Second)

    go func() {
        msgC := log.Subscribe(10)
        for {
            msg := <- msgC
            fmt.Printf("level: %s, time: %s, message: %s\n",
                msg.Level, msg.Timestamp.Format("15:04:05"), msg.Message)
        }
    }()
    log.Info("startup - version: 0.1")
    time.Sleep(10 * time.Millisecond)

    log.Debug("change to debug level")
    log.SetLevel(plog.Debug)
    log.Debug("level changed")

    // wait for the go routine above to consume the message
    time.Sleep(100 * time.Millisecond)
}
