package bluetooth

import (
	"testing"
	"log"
)
type mockNotifier struct {
	written []byte
	check bool
}

func (m *mockNotifier) Write(data []byte) (int, error) {
	m.written = data
	return 0, nil
}
func (m *mockNotifier) Done() bool {
	if m.check == false {
		m.check = true
		return false
	}
	return m.check 
}
func (m mockNotifier) Cap() int {
	return 0
}
func TestNotify(t *testing.T) {
	notifier := new(mockNotifier)
	go notify(notifier)
	
	log.Println("Sending message")
	dsStatus <- true
	
	if len(notifier.written) == 0 {
		t.Error("No message received through the notify method")
	} else {
		log.Printf("Message received: %s", string(notifier.written))
	}
}