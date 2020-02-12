package plog

import (
	"testing"
)

func TestGlobalLoggerIsAsSingleton(t *testing.T) {
	l1 := GlobalLogger()
	l2 := GlobalLogger()

	if l1 != l2 {
		t.Errorf("logger was not a singleton")
	}
}


func TestGlobalLogger(t *testing.T) {
	DropUnhandledMessages()
	l1 := GlobalLogger()
	l2 := GlobalLogger()

	// no logger attached - this message is lost
	l2.Info("i'm lost!")

	// init a stream logger
	stream := NewStreamLogger()
	msgC := stream.Subscribe(10)

	// attach the stream logger to 'l1'
	l1.Add(stream)

	// log with 'l2'
	l2.Info("did you see me?")

	// verify
	msg := <-msgC
	if msg.Message != "did you see me?" {
		t.Errorf("expected message not received!")
	}
}
