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

func (u *UserDisplay) SetGPIO(gpio Gpio) {
	u.gpio = gpio
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

func (u *UserDisplay) Error() {
	setHigh(u, RedLED)
}

func (u *UserDisplay) Reset() {
	setLow(u, YellowLED)
	setLow(u, GreenLED)
	setLow(u, RedLED)
}

func (u *UserDisplay) Open() {
	u.gpio.Open()
        setPinToOutput(u)
}

func setPinToOutput(u *UserDisplay) {
        u.gpio.Pin(GreenLED).Output()
        u.gpio.Pin(RedLED).Output()
        u.gpio.Pin(YellowLED).Output()
}

func (u *UserDisplay) Close() {
	u.gpio.Close()
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