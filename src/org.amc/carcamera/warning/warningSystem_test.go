package warning

import (
	"github.com/stianeikeland/go-rpio"
	"testing"
)

func TestReset(t *testing.T) {
	warning := UserDisplay{
		gpio: new(MockGpio),
	}

	warning.gpio.Pin(GreenLED).High()
	warning.gpio.Pin(RedLED).High()

	warning.Reset()

	if warning.gpio.Pin(GreenLED).Read() == rpio.Low &&
		warning.gpio.Pin(RedLED).Read() == rpio.Low {

	} else {

		t.Error("Call to Reset didn't reset lights")
	}
}

func TestOk(t *testing.T) {
	warning := UserDisplay{
		gpio: new(MockGpio),
	}

	warning.Ok()

	if warning.gpio.Pin(GreenLED).Read() != rpio.High {
		t.Errorf("Pin %d not set high\n", GreenLED)
	}

	if warning.gpio.Pin(RedLED).Read() == rpio.High {
		t.Error("Yellow and Red light still on")
	}
}

func TestNotOk(t *testing.T) {
	warning := UserDisplay{
		gpio: new(MockGpio),
	}
	warning.Error()
	warning.Ok()

	if warning.gpio.Pin(GreenLED).Read() == rpio.High {
		t.Errorf("Green Led shouldnt be light when Red Led is light")
	}
}

func TestError(t *testing.T) {
	warning := UserDisplay{
		gpio: new(MockGpio),
	}

	warning.Error()

	if warning.gpio.Pin(RedLED).Read() == rpio.Low {
		t.Errorf("Red Led should light")
	}
}

func TestOpenAndClose(t *testing.T) {
	warning := UserDisplay{
		gpio: new(MockGpio),
	}

	warning.Open()

	if open == false {
		t.Error("Open method not called")
	}

	warning.Close()

	if open != false {
		t.Error("Close method not called")
	}
}
