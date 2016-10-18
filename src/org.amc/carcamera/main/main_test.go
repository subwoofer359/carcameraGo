package main

import (
    "testing"
    "org.amc/carcamera/warning"
    "github.com/stianeikeland/go-rpio"
    "errors"
	
)


var mockGPIO *warning.MockGpio
var testapp *app

func setup() {
	mockGPIO = warning.NewMockGPIO()
	
	context["WORKDIR"] = "/tmp" // set workdir to /tmp for testing
	
	testapp = new(app)
	testapp.lights.SetGPIO(mockGPIO)
	testapp.Init()
}

func TestCreateCameraCommand(t *testing.T) {
	setup()
	command := createWebCamCommand()
	
	if command == nil {
		t.Error("Command not created")
	}
}

func TestAppInit(t *testing.T) {
	setup()
	
	if !mockGPIO.IsOpen() {
		t.Error("GPIO should be set to open")
	}
	
	if mockGPIO.Pin(warning.RedLED).Read() != rpio.Low {
		t.Error("Red light should be off")
	}
	
	if mockGPIO.Pin(warning.YellowLED).Read() != rpio.Low {
		t.Error("Yellow light should be off")
	}
	
	if mockGPIO.Pin(warning.GreenLED).Read() != rpio.Low {
		t.Error("Green light should be off")
	}
	
	if testapp.WebCamApp == nil {
		t.Error("WebCamApp should not be nil")
	}
}

func TestInitStorageManager(t *testing.T) {
	setup()
	
	if err := testapp.InitStorageManager(); err != nil {
		t.Error("Call to InitStorageManager() shouldn't throw an exception")
	}

	if testapp.WebCamApp.storageManager.WorkDir == nil {
		t.Error("Work Directory for StorageManager not set")
	}	
}

func TestInitStorageManagerError(t *testing.T) {
	setup()

	testapp.WebCamApp.storageManager = new(mainMockStorageManager)

	testapp.InitStorageManager()
	
	if mockGPIO.Pin(warning.RedLED).Read() != rpio.High {
		t.Error("Red light not turned on")
	}
		
}



func TestStartError(t *testing.T) {
	setup()
	
	testapp.WebCamApp.command = "/bin/l"
	testapp.WebCamApp.storageManager = GetMockStorageManagerLS()
	
	testapp.InitStorageManager()
	
	if err := testapp.Start(); err == nil {
		t.Error("An error should have been thrown")
	}
	
	if mockGPIO.Pin(warning.RedLED).Read() != rpio.High {
		t.Error("Red led should be light as there is an error in execution")
	}
}

//mainMockStorageManager Mock StorageManager
type mainMockStorageManager struct {
	mockStorageManager		
}

func (m *mainMockStorageManager) Init() error {
	return errors.New("Test StorageManager init failed")
	return nil
}

func (m *mainMockStorageManager) GetNextFileName() string {
	return ""
}
