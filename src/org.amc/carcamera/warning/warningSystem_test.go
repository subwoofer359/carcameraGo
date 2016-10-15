package warning

import (
	"testing"
	"github.com/stianeikeland/go-rpio"
)

func TestReset(t *testing.T) {
	warning := UserDisplay {
		gpio: new(mockGpio),
	}
	
	warning.gpio.Pin(GreenLED).High()
	warning.gpio.Pin(YellowLED).High()
	warning.gpio.Pin(RedLED).High()
	
	warning.Reset()
	
	if warning.gpio.Pin(GreenLED).Read() == rpio.Low && warning.gpio.Pin(YellowLED).Read() == rpio.Low &&
		warning.gpio.Pin(RedLED).Read() == rpio.Low {
			
		} else {
			
			t.Error("Call to Reset didn't reset lights")
		}
}

func TestWarning(t *testing.T) {
	warning := UserDisplay {
		gpio: new(mockGpio),
	}
	warning.Ok()
	
	warning.Warn()
	
	if warning.gpio.Pin(GreenLED).Read() == rpio.High {
		t.Error("Green LED should be off")
	}
	
	if warning.gpio.Pin(YellowLED).Read() == rpio.Low {
		t.Errorf("Pin %d not set high\n", YellowLED)
	}
}

func TestOk(t *testing.T) {
	warning := UserDisplay {
		gpio: new(mockGpio),
	}
	
	warning.Ok()
	
	if warning.gpio.Pin(GreenLED).Read() != rpio.High {
		t.Errorf("Pin %d not set high\n", GreenLED)
	}
	
	if warning.gpio.Pin(YellowLED).Read() == rpio.High ||
		warning.gpio.Pin(RedLED).Read() == rpio.High {
			t.Error("Yellow and Red light still on")	
		}
}

func TestNotOk(t *testing.T) {
	warning := UserDisplay {
		gpio: new(mockGpio),
	}
	
	warning.Warn()
	warning.Ok()
	
	if warning.gpio.Pin(GreenLED).Read() == rpio.High {
		t.Errorf("Green Led shouldnt be light when Yellow and Red Led are light")
	}
}


func TestError(t *testing.T) {
	warning := UserDisplay {
		gpio: new(mockGpio),
	}
	
	warning.Error()
	
	if warning.gpio.Pin(RedLED).Read() == rpio.Low {
		t.Errorf("Red Led should light")
	}
}

func TestOpenAndClose(t *testing.T) {
	warning := UserDisplay {
		gpio: new(mockGpio),
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