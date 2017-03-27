package bluetooth

import (
	"testing"
	"log"
	"time"
)

var (
	service *dashCamBTService
	notifier *mockNotifier
	errorMsg string
)


func testErrorInit() {
	service = getTestDashCamBTService()
	notifier = new(mockNotifier)
	errorMsg = "There has been an error"
}

func TestSendError(t *testing.T) {
	testErrorInit()
	
	go notifyError(notifier, service)
	log.Println("Sending Error message")
	
	service.SendError(errorMsg)
	
	log.Println("Sent Error message")
	service.Update()
	
	if service.errorMsg != errorMsg {
		t.Error("Error message not sent")
	}
	
	time.Sleep(100 * time.Millisecond)
	if len(notifier.written) == 0 {
		t.Error("No message received through the notify method")
	} else {
		log.Printf("Message received: %s", string(notifier.written))
	}
}

/**
 * No notification should be sent if error message is empty
 */
func TestNoSendNoError(t *testing.T) {
	testErrorInit()
	
	go notifyError(notifier, service)
	
	time.Sleep(100 * time.Millisecond)
	if notifier.wasWritten {
		t.Error("There was no message to send for notification")
	} 
}

/**
 * create mock where done always returns true
 */
type mockErrorNotifier struct {
	mockNotifier
}

func (m mockErrorNotifier) Done() bool {
	return false
}

func TestSendErrorOnce(t *testing.T) {
	testErrorInit()
	
	n := new(mockErrorNotifier)
	
	go notifyError(n, service)
	
	service.SendError(errorMsg)
	
	service.Update()
	
	if service.errorMsg != errorMsg {
		t.Error("Error message not sent")
	}
	
	time.Sleep(100 * time.Millisecond)
	if len(n.written) == 0 {
		t.Error("No message received through the notify method")
	} else {
		log.Printf("Message received: %s", string(n.written))
	}
	n.wasWritten = false
	
	time.Sleep(100 * time.Millisecond)
	
	if n.wasWritten {
		t.Error("Message should only be set once")
	}
}