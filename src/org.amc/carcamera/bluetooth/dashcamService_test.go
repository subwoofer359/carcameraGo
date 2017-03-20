package bluetooth

import (
	"testing"
	"log"
	"time"
)
type mockNotifier struct {
	written []byte
	check bool
}

func (m *mockNotifier) Write(data []byte) (int, error) {
	m.written = data
	m.check = true
	return 0, nil
}
func (m *mockNotifier) Done() bool {
	return m.check 
}
func (m mockNotifier) Cap() int {
	return 0
}
func TestNotify(t *testing.T) {
	notifier := new(mockNotifier)
	go notify(notifier)
	log.Println("Sending message true")
	dcBTServ.dsStatus <- true
	log.Println("Sending message false")
	dcBTServ.dsStatus <- false // send second message to cause block
	
	time.Sleep(100 * time.Millisecond)
	if len(notifier.written) == 0 {
		t.Error("No message received through the notify method")
	} else {
		log.Printf("Message received: %s", string(notifier.written))
	}
}