package plog


import "testing"
import "reflect"
import "time"

func TestBroadcast(t *testing.T) {

	stream1 := NewStreamLogger()
	msgC1 := stream1.Subscribe(10)

	stream2 := NewStreamLogger()
	msgC2 := stream2.Subscribe(10)

	//
	log := NewBroadcastLogger(stream1, stream2)
	log.Info("info")
	log.Debug("debug")
	log.Warn("warn")


	// check
	lock := make(chan bool)
	go func() {
		msgs1 := slurp(msgC1)
		msgs2 := slurp(msgC2)

		if ! reflect.DeepEqual(msgs1, msgs2) {
			t.Errorf("msgs1 != msgs2 - msgs1: %v, msgs2: %v", msgs1, msgs2)
		}
		lock <- true
	}()

	stream1.WaitForSubscribers(100 * time.Millisecond)
	stream2.WaitForSubscribers(100 * time.Millisecond)
	<-lock

}


func slurp(s <-chan LogMessage) []LogMessage {
	var msgs []LogMessage
	for msg := range s {
		msgs = append(msgs, msg)
	}
	return msgs
}
