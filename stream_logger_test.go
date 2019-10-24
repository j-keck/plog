package plog

import (
	"testing"
	"strings"
	"io/ioutil"
	"time"
	"bytes"
	"sync/atomic"
)


func TestOneSubscriber(t *testing.T) {
	log := NewStreamLogger()
	log.SetStderr(ioutil.Discard)

	sub := log.Subscribe(1)

	message := "test message"
	log.Info(message)
	msg := <-sub

	if msg.Message != message {
		t.Errorf("%s != %s", message, msg.Message)
	}
}


func TestTwoSubscribers(t *testing.T) {
	log := NewStreamLogger()
	log.SetStderr(ioutil.Discard)

	sub1 := log.Subscribe(1)
	sub2 := log.Subscribe(1)

	message := "test message"
	log.Info(message)
	msg1 := <-sub1
	msg2 := <-sub2

	if msg1.Message != message || msg2.Message != message {
		t.Errorf("%s != %s || %s != %s",
			message, msg1.Message,
		    message, msg2.Message)
	}
}


func TestFullChannelShouldNotBlock(t *testing.T) {
	log := NewStreamLogger()
	log.SetStderr(ioutil.Discard)

	// unbuffered channel
	log.Subscribe(0)

	// this would block, if the logger doesn't
	// handle full channels
	log.Info("test message")
}


func TestWaitForSubscribers(t *testing.T) {
	log := NewStreamLogger()
	log.SetStderr(ioutil.Discard)

	sub := log.Subscribe(1)
	log.Info("test message")

	var counter int32
	go func() {
		log.WaitForSubscribers(5 * time.Second)
		if atomic.LoadInt32(&counter) == 0 {
			t.Error("message not received")
		}
	}()

	time.Sleep(500 * time.Millisecond)
	<-sub
	atomic.AddInt32(&counter, 1)
}

func TestWaitForSubscribersShouldTimeout(t *testing.T) {
	log := NewStreamLogger()

	var stderr bytes.Buffer
	log.SetStderr(&stderr)

	log.Subscribe(1)
	log.Info("test message")

	log.WaitForSubscribers(100 * time.Millisecond)
	if !strings.Contains(stderr.String(), "timeout after") {
		t.Error("timeout error message on stderr expected")
	}
}
