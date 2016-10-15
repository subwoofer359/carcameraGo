package main

import (
    "testing"
    "org.amc/carcamera/warning"
    "github.com/stianeikeland/go-rpio"
    "errors"
	
)


func TestCreateCameraCommand(t *testing.T) {
	command := createWebCamCommand()
	
	if command == nil {
		t.Error("Command not created")
	}
}

func TestAppInit(t *testing.T) {
	var myapp app;
	mockGPIO := warning.NewMockGPIO()
	myapp.lights.SetGPIO(mockGPIO)
	myapp.Init()
	
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
	
	if myapp.WebCamApp == nil {
		t.Error("WebCamApp should not be nil")
	}
}

func TestInitStorageManager(t *testing.T) {
	
	var myapp app
	myapp.lights.SetGPIO(warning.NewMockGPIO())
	
	myapp.Init()
	
	if err := myapp.InitStorageManager(); err != nil {
		t.Error("Call to InitStorageManager() shouldn't thrown an exception")
	}
	
	
	
	if myapp.WebCamApp.storageManager.WorkDir == nil {
		t.Error("Work Directory for StorageManager not set")
	}	
}

func TestInitStorageManagerError(t *testing.T) {
	var myapp app
	mockGPIO := warning.NewMockGPIO()
	myapp.lights.SetGPIO(mockGPIO)
	
	myapp.Init()

	myapp.WebCamApp.storageManager = new(mainMockStorageManager)

	myapp.InitStorageManager()
	
	if mockGPIO.Pin(warning.RedLED).Read() != rpio.High {
		t.Error("Red light not turned on")
	}
		
}



func TestStartError(t *testing.T) {
	var myapp app;
	mockGPIO := warning.NewMockGPIO()
	myapp.lights.SetGPIO(mockGPIO)
	myapp.Init()
	
	
	myapp.WebCamApp.command = "/bin/l"
	myapp.WebCamApp.storageManager = GetMockStorageManagerLS()
	
	myapp.InitStorageManager()
	
	if err := myapp.Start(); err == nil {
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
