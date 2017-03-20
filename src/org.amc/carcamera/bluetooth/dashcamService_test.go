package bluetooth

import (
	"testing"
	"log"
	"time"
)

/*
 * A empty dashCamBTService for testing purposes
 */
func getTestDashCamBTService() *dashCamBTService {
	service := new (dashCamBTService)
	service.dsStatus = make(chan bool, 1)
	return service
}

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
	
	service := getTestDashCamBTService()
	notifier := new(mockNotifier)
	go notify(notifier, service)
	log.Println("Sending message true")
	service.SendStatus(true)
	service.Update()
	log.Println("Sending message false")
	service.SendStatus(false) // send second message to cause block
	service.Update()
	
	time.Sleep(100 * time.Millisecond)
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
	
	
	
	if service.statusChanged == true {
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
	
	if service.statusChanged != true {
		t.Error("statusChanged should be true")
	}
	
	if service.status == previous {
		t.Error("Should be an update of status")
	} 
}

func TestSetDifferentStatus(t *testing.T) {
	service := getTestDashCamBTService()
	service.setStatus(true)
	
	
	if service.status != true {
		t.Error("Boolean value for status wasn't set")
	}
	
	if service.statusChanged != true {
		t.Error("Change of status value not detected")
	}
}

func TestSetSameStatus(t *testing.T) {
	service := getTestDashCamBTService()
	
	
	if service.status != false {
		t.Error("Boolean value for status wasn't set")
	}
	
	if service.statusChanged == true {
		t.Error("Change of status value incorrectly detected")
	}
}