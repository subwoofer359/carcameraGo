package warning

import (
	"testing"
)

func TestReset(t *testing.T) {
	warning := UserDisplay{
		gpio: new(MockGpio),
	}

	warning.gpio.Pin(GreenLED).High()
	warning.gpio.Pin(RedLED).High()

	warning.Reset()

	if warning.gpio.Pin(GreenLED).Read() == Low &&
		warning.gpio.Pin(RedLED).Read() == Low {

	} else {

		t.Error("Call to Reset didn't reset lights")
	}
}

func TestOk(t *testing.T) {
	warning := UserDisplay{
		gpio: new(MockGpio),
	}

	warning.Ok()

	if warning.gpio.Pin(GreenLED).Read() != High {
		t.Errorf("Pin %d not set high\n", GreenLED)
	}

	if warning.gpio.Pin(RedLED).Read() == High {
		t.Error("Yellow and Red light still on")
	}
}

func TestNotOk(t *testing.T) {
	warning := UserDisplay{
		gpio: new(MockGpio),
	}
	warning.Error()
	warning.Ok()

	if warning.gpio.Pin(GreenLED).Read() == High {
		t.Errorf("Green Led shouldnt be light when Red Led is light")
	}
}

func TestError(t *testing.T) {
	warning := UserDisplay{
		gpio: new(MockGpio),
	}

	warning.Error()

	if warning.gpio.Pin(RedLED).Read() == Low {
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

func TestPowerError(t *testing.T) {
	warning := UserDisplay{
		gpio: new(MockGpio),
	}

	warning.Ok()

	warning.PowerError()

	if warning.gpio.Pin(GreenLED).Read() == High {
		t.Errorf("Green LED should be off")
	}
}
