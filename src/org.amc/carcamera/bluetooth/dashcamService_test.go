package bluetooth

import (
	"log"
	"sync"
	"testing"
)

/*
 * A empty dashCamBTService for testing purposes
 */
func getTestDashCamBTService() *dashCamBTService {
	service := new(dashCamBTService)
	service.dsStatus = make(chan bool, 1)
	service.dsErrorMsg = make(chan string, 1)
	return service
}

type mockNotifier struct {
	wasWritten bool
	written    []byte
	check      bool
	wg         *sync.WaitGroup
}

func (m *mockNotifier) Write(data []byte) (int, error) {
	defer m.wg.Done()
	m.written = data
	m.wasWritten = true
	m.check = true
	return len(data), nil
}

func (m *mockNotifier) Done() bool {
	return m.check
}

func (m mockNotifier) Cap() int {
	return 0
}

func TestNotify(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(1)
	service := getTestDashCamBTService()
	notifier := mockNotifier{
		wg: &wg,
	}

	go notifyStatus(&notifier, service)

	log.Println("Sending message true")
	service.SendStatus(true)
	service.Update()
	wg.Wait()

	if len(notifier.written) == 0 {
		t.Error("No message received through the notify method")
	} else {
		log.Printf("Message received: %s", string(notifier.written))
	}
}

func TestUpdateNoUpdate(t *testing.T) {
	service := getTestDashCamBTService()

	previous := service.status

	service.Update()

	if service.statusChanged {
		t.Error("statusChanged should be false")
	}

	if service.status != previous {
		t.Error("Should be no update of status")
	}
}

func TestUpdateUpdate(t *testing.T) {
	service := getTestDashCamBTService()

	previous := service.status
	//sending boolean
	service.dsStatus <- true

	service.Update()

	if !service.statusChanged {
		t.Error("statusChanged should be true")
	}

	if service.status == previous {
		t.Error("Should be an update of status")
	}
}

func TestSetDifferentStatus(t *testing.T) {
	service := getTestDashCamBTService()
	service.setStatus(true)

	if !service.status {
		t.Error("Boolean value for status wasn't set")
	}

	if !service.statusChanged {
		t.Error("Change of status value not detected")
	}
}

func TestSetSameStatus(t *testing.T) {
	service := getTestDashCamBTService()

	if service.status {
		t.Error("Boolean value for status wasn't set")
	}

	if service.statusChanged {
		t.Error("Change of status value incorrectly detected")
	}
}
