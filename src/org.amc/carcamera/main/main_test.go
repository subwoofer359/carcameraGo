package main

import (
	"errors"
	"github.com/stianeikeland/go-rpio"
	"github.com/stretchr/testify/assert"
	"log"
	C "org.amc/carcamera/constants"
	"org.amc/carcamera/userupdate"
	"org.amc/carcamera/warning"
	"os"
	"sort"
	"testing"
)

var mockGPIO *warning.MockGpio
var testapp *app

func setup() {
	mockGPIO = warning.NewMockGPIO()

	context = make(map[string]interface{})

	context[C.COMMAND] = "/bin/ls"
	context[C.WORKDIR] = "/tmp" // set workdir to /tmp for testing
	context[C.PREFIX] = "video"
	context[C.SUFFIX] = ".h264"
	context[C.OPTIONS] = []string{"-ss", "70000", "-rot", "90"}
	context[C.TIMEOUT] = "1000"

	log.Println(context)
	testapp = new(app)

	ledService := new(userupdate.LEDService)
	ledService.SetGPIO(mockGPIO)

	testapp.Init()

	testapp.message.AddService(ledService)

	testapp.message.Init()
}

func TestCreateCameraCommand(t *testing.T) {

	setup()
	command := createWebCamCommand()

	if command == nil {
		t.Error("Command not created")
	}

	if sort.SearchStrings(command.args, "-rot") >= len(command.args) {
		t.Error("Option Rotate not found")
	}

	if sort.SearchStrings(command.args, "90") >= len(command.args) {
		t.Error("Option Rotate argument not found")
	}

	if sort.SearchStrings(command.args, "help") < len(command.args) {
		t.Error("Unspecified option found")
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

	if mockGPIO.Pin(warning.GreenLED).Read() != rpio.Low {
		t.Error("Green light should be off")
	}

	if testapp.WebCamApp == nil {
		t.Error("WebCamApp should not be nil")
	}
}

func TestInitStorageManager(t *testing.T) {
	setup()

	testapp.WebCamApp.storageManager = GetMockStorageManagerLS()

	testapp.InitStorageManager()

	assert.Equal(t, testapp.WebCamApp.storageManager.WorkDir(), context[C.WORKDIR],
		"Work directory not set properly")
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

func TestLoadConfiguration(t *testing.T) {
	testapp = new(app)
	var filename string = os.ExpandEnv("$GOPATH") + "/configuration.json"
	err := testapp.LoadConfiguration(filename)

	if err != nil {
		t.Error(err)
	}

	if context[C.COMMAND] == "" {
		t.Errorf("%s shouldn't be empty\n", C.COMMAND)
	}

}

//mainMockStorageManager Mock StorageManager
type mainMockStorageManager struct {
	mockStorageManager
}

func (m *mainMockStorageManager) Init() error {
	return errors.New("Test StorageManager init failed")
}

func (m *mainMockStorageManager) GetNextFileName() string {
	return ""
}
