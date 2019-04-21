package main

import "github.com/j-keck/plog"
import "fmt"
import "time"

func main() {
    log := plog.NewStreamLogger()

    go func() {
        logC := log.Subscribe(10)
        for {
            msg := <- logC
            fmt.Printf("level: %s, message: %s\n",
                msg.Level, msg.Message)
        }
    }()

    log.Info("startup")
    log.Info("shutdown")
    log.WaitForSubscribers(100 * time.Millisecond)
}
