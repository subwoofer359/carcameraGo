package warning

import (
	"github.com/stianeikeland/go-rpio"
	//"log"	
)

const (
	GreenLED int = 24

	YellowLED int = 23

	RedLED int = 18
)



type UserDisplay struct {
	gpio Gpio
}

func (u *UserDisplay) Warn() {
	setLow(u, GreenLED)
	setHigh(u, YellowLED)
}

func (u *UserDisplay) Ok() {
	if !(isHigh(u, YellowLED) || isHigh(u, RedLED)) {
			setHigh(u, GreenLED)
	}
}	

func (u *UserDisplay) Reset() {
	setLow(u, YellowLED)
	setLow(u, GreenLED)
	setLow(u, RedLED)
}

func setHigh(u *UserDisplay, colour int) {
	u.gpio.Pin(colour).High()
}

func setLow(u *UserDisplay, colour int) {
	u.gpio.Pin(colour).Low()
}

func read(u *UserDisplay, colour int) rpio.State {
	return u.gpio.Pin(colour).Read()
}

func isHigh(u *UserDisplay, colour int) bool {
	return read(u, colour) == rpio.High
}