package userupdate

import (
	"testing"

	"org.amc/carcamera/warning"
)

func TestLEDServiceInit(t *testing.T) {
	ledService := new(LEDService)

	mockGPIO := warning.NewMockGPIO()

	ledService.SetGPIO(mockGPIO)

	err := ledService.Init()

	if err != nil {
		t.Error(err)
	}

	if !mockGPIO.IsOpen() {
		t.Error("GPIO wasn't opened")
	}
}

func TestLEDServiceNotStarted(t *testing.T) {
	ledService := new(LEDService)

	mockGPIO := warning.NewMockGPIO()

	ledService.SetGPIO(mockGPIO)

	err := ledService.Init()

	if err != nil {
		t.Error(err)
	}

	if !mockGPIO.IsOpen() {
		t.Error("GPIO wasn't opened")
	}

	greenPin := mockGPIO.Pin(warning.GreenLED)

	if greenPin.Read() != warning.Low {
		t.Error("Green light should be off")
	}

	redPin := mockGPIO.Pin(warning.RedLED)

	if redPin.Read() != warning.Low {
		t.Error("Red light should be off")
	}
}

func TestLEDServiceStarted(t *testing.T) {
	ledService := new(LEDService)

	mockGPIO := warning.NewMockGPIO()

	ledService.SetGPIO(mockGPIO)

	err := ledService.Init()

	if err != nil {
		t.Error(err)
	}

	if !mockGPIO.IsOpen() {
		t.Error("GPIO wasn't opened")
	}

	ledService.Started()

	greenPin := mockGPIO.Pin(warning.GreenLED)

	if greenPin.Read() != warning.High {
		t.Error("Green light should be on")
	}

	redPin := mockGPIO.Pin(warning.RedLED)

	if redPin.Read() != warning.Low {
		t.Error("Red light should be off")
	}
}

func TestLEDServiceError(t *testing.T) {
	ledService := new(LEDService)

	mockGPIO := warning.NewMockGPIO()

	ledService.SetGPIO(mockGPIO)

	err := ledService.Init()

	if err != nil {
		t.Error(err)
	}

	if !mockGPIO.IsOpen() {
		t.Error("GPIO wasn't opened")
	}

	ledService.Started()

	ledService.Error("Error")

	redPin := mockGPIO.Pin(warning.RedLED)

	if redPin.Read() != warning.High {
		t.Error("Red light should be on")
	}
}

func TestLEDServiceStopped(t *testing.T) {
	ledService := new(LEDService)

	mockGPIO := warning.NewMockGPIO()

	ledService.SetGPIO(mockGPIO)

	ledService.Init()

	ledService.Started()

	redPin := mockGPIO.Pin(warning.RedLED)

	if redPin.Read() == warning.High {
		t.Error("Red light should not be on")
	}

	ledService.Stopped()

	if redPin.Read() != warning.High {
		t.Error("Red light should be on")
	}
}

func TestLEDServiceIsClosed(t *testing.T) {
	ledService := new(LEDService)

	mockGPIO := warning.NewMockGPIO()

	ledService.SetGPIO(mockGPIO)

	ledService.Init()

	ledService.Close()

	if mockGPIO.IsOpen() {
		t.Error("GPIO wasn't closed")
	}
}
