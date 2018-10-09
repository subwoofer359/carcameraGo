package warning

import (
	"github.com/stianeikeland/go-rpio"
)

// RpioImpl wrapper for the rpio package
//
type RpioImpl struct{}

func (r RpioImpl) Open() error {
	return rpio.Open()
}
func (r RpioImpl) Close() error {
	return rpio.Close()
}

func (r RpioImpl) Pin(number int) GpioPin {
	rpin := rpio.Pin(number)
	return &rpiopinWrapper{
		pin: rpin,
	}
}

type rpiopinWrapper struct {
	pin rpio.Pin
}

func (r *rpiopinWrapper) High() {
	r.pin.High()
}

//Wrapper for rpio pin
func (r *rpiopinWrapper) Low() {
	r.pin.Low()
}

func (r *rpiopinWrapper) Input() {
	r.pin.Input()
}

func (r *rpiopinWrapper) Output() {
	r.pin.Output()
}

func (r *rpiopinWrapper) Read() State {
	switch rpio.ReadPin(r.pin) {
	case rpio.High:
		return High
	case rpio.Low:
		return Low
	default:
		return Low
	}
}
