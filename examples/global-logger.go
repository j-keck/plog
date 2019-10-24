package main


import "github.com/j-keck/plog"

func main() {
  // initialize the logger
  log := plog.GlobalLogger().Add(plog.NewDefaultConsoleLogger())
  log.Info("in 'main'")

  other()
}

func other() {
  log := plog.GlobalLogger()
  log.Info("in 'other'")
}
