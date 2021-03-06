package main

import (
	"errors"
	"log"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	C "org.amc/carcamera/constants"
	"org.amc/carcamera/powercontrol"
	"org.amc/carcamera/runner"
	"org.amc/carcamera/userupdate"
	"org.amc/carcamera/warning"
)

var mockGPIO *warning.MockGpio
var testapp *app

func contextTestSetup() {
	context = make(map[string]interface{})

	context[C.COMMAND] = "/bin/l"
	context[C.WORKDIR] = "/tmp" // set workdir to /tmp for testing
	context[C.PREFIX] = "video"
	context[C.SUFFIX] = ".h264"
	context[C.OPTIONS] = []string{"-ss", "70000", "-rot", "90"}
	context[C.TIMEOUT] = "1000"
}

func setup() {
	//disable shutdown command
	mockEndCommand := &mockEndCmd{
		shutdownCMD: []string{"ls", "-l", "-a"},
		syncCMD:     []string{"ls"},
	}

	mockGPIO = warning.NewMockGPIO()

	contextTestSetup()

	log.Println(context)

	testapp = new(app)

	pc := new(powercontrol.MockPowerControl)

	pc.Init()

	testapp.endCmd = mockEndCommand

	testapp.powerControl = pc.PowerOff()

	pc.Start()

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

	if sort.SearchStrings(command.Args(), "-rot") >= len(command.Args()) {
		t.Error("Option Rotate not found")
	}

	if sort.SearchStrings(command.Args(), "90") >= len(command.Args()) {
		t.Error("Option Rotate argument not found")
	}

	if sort.SearchStrings(command.Args(), "help") < len(command.Args()) {
		t.Error("Unspecified option found")
	}

}

func TestAppInit(t *testing.T) {
	setup()

	if !mockGPIO.IsOpen() {
		t.Error("GPIO should be set to open")
	}

	if mockGPIO.Pin(warning.RedLED).Read() != warning.Low {
		t.Error("Red light should be off")
	}

	if mockGPIO.Pin(warning.GreenLED).Read() != warning.Low {
		t.Error("Green light should be off")
	}

	if testapp.WebCamApp == nil {
		t.Error("WebCamApp should not be nil")
	}

	if testapp.runnerFactory != defaultFactory {
		t.Error("WebCamApp runnerFactory not set to default")
	}

	if testapp.powerControl == nil {

		t.Error("PowerControl not set up")
	}
}

func TestAppInitNewFactory(t *testing.T) {
	contextTestSetup()
	testapp = new(app)

	myMockRunner := new(mockRunnerFactory)

	assert.Nil(t, testapp.runnerFactory)

	testapp.runnerFactory = myMockRunner

	testapp.Init()

	assert.NotEqual(t, myMockRunner, defaultFactory)

	assert.Equal(t, myMockRunner, testapp.runnerFactory)
}

// Test stopped error
var errTestStopped = errors.New("Test Stopped")

func TestInitStorageManager(t *testing.T) {
	setup()

	testapp.WebCamApp.SetStorageManager(runner.GetMockStorageManagerLS(context))

	testapp.InitStorageManager()

	assert.Equal(t, testapp.WebCamApp.StorageManager().WorkDir(), context[C.WORKDIR],
		"Work directory not set properly")
}

func TestInitStorageManagerError(t *testing.T) {
	setup()

	testapp.WebCamApp.SetStorageManager(new(runner.MainMockStorageManager))

	testapp.InitStorageManager()

	if mockGPIO.Pin(warning.RedLED).Read() != warning.High {
		t.Error("Red light not turned on")
	}

}

func TestStartError(t *testing.T) {
	setup()

	testapp.WebCamApp.SetCommand("/bin/l")
	testapp.WebCamApp.SetStorageManager(runner.GetMockStorageManagerLS(context))

	testapp.InitStorageManager()

	if err := testapp.Start(); err == nil {
		t.Error("An error should have been thrown")
	}

	if mockGPIO.Pin(warning.RedLED).Read() != warning.High {
		t.Error("Red led should be light as there is an error in execution")
	}
}

func TestStart(t *testing.T) {
	setup()

	testapp.runnerFactory = new(mockRunnerFactory)

	//Set up test time out
	testTimeout := 5 * time.Second

	timeoutChan := make(<-chan time.Time)

	timeoutChan = time.After(testTimeout)

	result := make(chan error)

	defer close(result)

	go func() {
		result <- testapp.Start()
	}()

	select {
	case <-timeoutChan:
		t.Fatal("Test timed out")
	case err := <-result:
		checkRunnerReturn(t, err)
	}
}

func checkRunnerReturn(t *testing.T, err error) {
	switch err {
	case nil:
		break
	case powercontrol.ErrPowerFault:
		log.Println("TestStartPowerOff: Power fault error returned ")
		break
	case errTestStopped:
		log.Println("TestStartPowerOff: test Stopped ")
		break
	default:
		t.Error(err)
	}
}

func TestStartPowerOff(t *testing.T) {
	setup()

	testapp.runnerFactory = new(slowMockRunnerFactory)

	newPowerControl := new(powercontrol.PPowerControl)

	newPowerControl.Init()

	testapp.powerControl = newPowerControl.PowerOff()

	//Set up test time out
	testTimeout := 10 * time.Second

	timeoutChan := make(<-chan time.Time)

	timeoutChan = time.After(testTimeout)

	result := make(chan error)

	defer close(result)

	go func() {
		result <- testapp.Start()
	}()

	newPowerControl.Start()
	select {
	case <-timeoutChan:
		t.Fatal("Test timed out")
	case err := <-result:
		if err == errTestStopped {
			t.Fatal(err)
		}
		checkRunnerReturn(t, err)
	}
}
