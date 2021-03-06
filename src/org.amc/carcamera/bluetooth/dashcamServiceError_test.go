package bluetooth

import (
	"log"
	"sync"
	"testing"
	"time"
)

var (
	service  *dashCamBTService
	notifier *mockNotifier
	errorMsg string
)

func testErrorInit() {
	service = getTestDashCamBTService()
	notifier = new(mockNotifier)
	errorMsg = "There has been an error"
}

func TestSendError(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(1)
	// reduce notify delay
	oldNotifyDelay := NOTIFY_DELAY
	NOTIFY_DELAY = 0

	testErrorInit()
	notifier.wg = &wg

	go notifyError(notifier, service)
	log.Println("Sending Error message")

	service.SendError(errorMsg)

	log.Println("Sent Error message")
	service.Update()

	if service.errorMsg != errorMsg {
		t.Error("Error message not sent")
	}

	wg.Wait()
	if len(notifier.written) == 0 {
		t.Error("No message received through the notify method")
	} else {
		log.Printf("Message received: %s", string(notifier.written))
	}

	//restore notify delay
	NOTIFY_DELAY = oldNotifyDelay
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
	var wg sync.WaitGroup
	wg.Add(1)

	testErrorInit()

	n := new(mockErrorNotifier)
	n.wg = &wg

	go notifyError(n, service)

	service.SendError(errorMsg)

	service.Update()

	if service.errorMsg != errorMsg {
		t.Error("Error message not sent")
	}

	wg.Wait()
	if !n.wasWritten {
		t.Error("No message received through the notify method")
	} else {
		log.Printf("Message received: %s", string(n.written))
	}

	n.wasWritten = false

	time.Sleep(1000 * time.Millisecond)

	if n.wasWritten {
		t.Error("Message should only be set once")
	}
}
