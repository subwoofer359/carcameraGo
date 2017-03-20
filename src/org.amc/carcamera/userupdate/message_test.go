package userupdate

import (
    "testing"
    "errors"
)

type testService struct {
	initCalled bool
	started bool
	errMessage string
}

func (testS *testService) Init() error {
	testS.initCalled = true
	return nil
}

func (testS *testService) Started() {
	testS.started = true
}

func (testS *testService) Error(message string) {
	testS.errMessage = message
}



type testServiceFail struct {
	testService
}

func (testS *testServiceFail) Init() error {
	
	return errors.New("Test Error")
}


func TestMessageAddService(t *testing.T) {
	message := new(Message)
	tService := new(testService)
	
	if len(message.services) != 0 {
		t.Error("The list of services should be empty")
	}
	
	message.AddService(tService)
	
	if len(message.services) != 1 {
		t.Error("The list of services should contain only one service")
	}
}

func TestMessageInit(t *testing.T) {
	message := new(Message)
	tService := new(testService)
	message.AddService(tService)
	
	message.Init()
	
	if tService.initCalled != true {
		t.Error("Init should have been called on the service")
	}
}

func TestMessageInitThrowsError(t *testing.T) {
	message := new(Message)
	tService := new(testServiceFail)
	message.AddService(tService)
	
	err := message.Init()
	
	if err == nil {
		t.Error("Error should have been thrown on Init() call")
	}
}

func TestSendErrorMessage(t *testing.T) {
	message := new(Message)
	tService := new(testService)
	errorMsg := "error!!!"
	
	message.AddService(tService)
	
	message.Error(errorMsg)
	
	if tService.errMessage != errorMsg {
		t.Error("Error message not passed to service")
	}
}

func TestStarted(t *testing.T) {
	message := new(Message)
	tService := new(testService)
	message.AddService(tService)
	
	message.Started()
	
	if tService.started == false {
		t.Error("Service not started")
	}
}