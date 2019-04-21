package plog

import (
	"os"
	"io"
	"time"
	"fmt"
)


type streamLogger struct {
	logger
	subscribers []chan LogMessage
	stderr      io.Writer
}

func NewStreamLogger() *streamLogger {
	self := new(streamLogger)
	self.SetLevel(Info)
	self.SetStderr(os.Stderr)

	self.logImpl = self

	return self
}


// SetStderr to overwrite 'stderr'
func (self *streamLogger) SetStderr(w io.Writer) {
	self.stderr = w
}

// Subscribe to log messages
func (self *streamLogger) Subscribe(bufferSize int) <-chan LogMessage {
	sub := make(chan LogMessage, bufferSize)
	self.subscribers = append(self.subscribers, sub)
	return sub
}

func minDuration(a, b time.Duration) time.Duration {
	if a < b {
		return a
	}
	return b
}
func (self *streamLogger) WaitForSubscribers(timeout time.Duration) {
	for _, c := range self.subscribers {
		const attempts = 10
		checkDelay := minDuration(timeout / attempts, 500 * time.Millisecond)

		for i := 0; ; i++ {
			if i >= attempts {
				timeoutMsg := fmt.Sprintf("plog.WaitForSubscribers timeout after %s - ignore Subscriber", timeout.String())
				logMsg := newLogMessage(Warn, timeoutMsg)
				fmt.Fprintln(self.stderr, logMsg.String())
				break
			}

			if len(c) == 0 {
				close(c)
				break
			}

			time.Sleep(checkDelay)
		}
	}
}


func (self *streamLogger) log(msg LogMessage) {
	if msg.Level < self.level {
		return
	}

	if len(self.subscribers) > 0 {
		for _, sub := range self.subscribers {
			if len(sub) < cap(sub) {
				sub <- msg
			} else {
				overflowMessage := newLogMessage(Warn, "channel full - discard msg: "+msg.String())
				fmt.Fprintln(self.stderr, overflowMessage.String())
			}
		}
	}
}
