package main


import "github.com/j-keck/plog"
import "fmt"
import "time"

func main() {
    log := plog.NewStreamLogger()
    logC := log.Subscribe(10)

    log.Info("startup")
    log.Debug("change to debug level")
    log.SetLevel(plog.Debug)
    log.Debug("level changed")
    log.Infof("2 + 2 = %d", 2 + 2)

    go func() {
      for msg := range logC {
        fmt.Printf("%s: %s\n", msg.Level, msg.Message)
      }
    }()

    log.WaitForSubscribers(100 * time.Millisecond)
}
