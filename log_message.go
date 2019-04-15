package plog

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogMessage struct {
	Level     LogLevel
	Timestamp time.Time
	File      string
	Line      int
	Message   string
}

func (self *LogMessage) String() string {
	return fmt.Sprintf("%s | %s", self.Level, self.Message)
}

func newLogMessage(level LogLevel, message string) LogMessage {
	logMessage := LogMessage{Timestamp: time.Now(), Level: level, Message: message}

	if _, file, line, ok := runtime.Caller(2); ok {
		fileName := path.Base(file)
		logMessage.File = strings.TrimSuffix(fileName, path.Ext(fileName))
		logMessage.Line = line
	}

	return logMessage
}