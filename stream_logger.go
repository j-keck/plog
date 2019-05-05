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


// NewStreamLogger creates a new logger, where the log messages
// are emitted over a go chan. Use 'Subscribe(int)' get a go chan.
func NewStreamLogger() *streamLogger {
	self := new(streamLogger)
	initLogger(self)
	self.SetStderr(os.Stderr)
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

// WaitForSubscribers waits till all consumers have received all
// messages.
//
// Note: This function does NOT block until the consumers have
//       PROCESSED the message (receiving != processing)
//
// TODO: should this return
//   - an error on timeout
//   - return some stats
func (self *streamLogger) WaitForSubscribers(timeout time.Duration) {
	minDuration := func (a, b time.Duration) time.Duration {
		if a < b {
			return a
		}
		return b
	}


	// check repeatedly if all consumers have received all messages.
	for _, c := range self.subscribers {
		const attempts = 10
		checkDelay := minDuration(timeout / attempts, 100 * time.Millisecond)

		for i := 0; ; i++ {
			if i >= attempts {

				logMsg := newLogMessage(
					Warn,
					fmt.Sprintf("plog.WaitForSubscribers timeout after %s - ignore Subscriber", timeout.String()),
				)

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
