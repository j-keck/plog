package plog

import "testing"
import "runtime"
import "path"
import "strings"
import "math"
import "io/ioutil"
import "time"
import "bytes"

func TestCallerLocation(t *testing.T) {
	log := NewConsoleLogger()
	log.SetStdout(ioutil.Discard)
	log.SetStderr(ioutil.Discard)
	stream := log.Subscribe(1)

	log.Info("ping")
	msg := <-stream

	_, file, line, _ := runtime.Caller(0)
	file = strings.TrimSuffix(path.Base(file), path.Ext(file))
	if file != msg.File || math.Abs(float64(line-msg.Line)) > 5 {
		t.Errorf("wrong caller location expected: %s:(%d+-5), received: %s:%d",
			file, line, msg.File, msg.Line)
	}
}

func TestMultiSubscriberShouldReceiveTheSameMsg(t *testing.T) {
	log := NewConsoleLogger()
	log.SetStdout(ioutil.Discard)
	log.SetStderr(ioutil.Discard)

	c1 := log.Subscribe(1)
	c2 := log.Subscribe(1)

	log.Info("<message>")

	msg1 := <-c1
	msg2 := <-c2

	if msg1.Message != "<message>" {
		t.Error("unexepcted log message")
	}

	if msg1 != msg2 {
		t.Error("mess<ages differ")
	}
}

func TestFullChannelShouldNotBlock(t *testing.T) {
	log := NewConsoleLogger()
	log.SetStdout(ioutil.Discard)
	log.SetStderr(ioutil.Discard)

	// unbuffered channel
	log.Subscribe(0)

	// this would block on a unbuffered channel
	log.Info("test message")
}

func TestWaitForSubscribers(t *testing.T) {
	log := NewConsoleLogger()
	log.SetStdout(ioutil.Discard)
	log.SetStderr(ioutil.Discard)

	sub := log.Subscribe(1)
	log.Info("test message")

	msgReceived := false
	go func() {
		time.Sleep(10 * time.Millisecond)
		<-sub
		msgReceived = true
	}()

	log.WaitForSubscribers(100 * time.Millisecond)

	if !msgReceived {
		t.Error("message not received")
	}
}

func TestWaitForSubscribersShouldTimeout(t *testing.T) {
	log := NewConsoleLogger()
	log.SetStdout(ioutil.Discard)
	var stderr bytes.Buffer
	log.SetStderr(&stderr)

	log.Subscribe(1)
	log.Info("test message")

	log.WaitForSubscribers(100 * time.Millisecond)
	if !strings.Contains(stderr.String(), "timeout after") {
		t.Error("timeout error message expected")
	}
}
