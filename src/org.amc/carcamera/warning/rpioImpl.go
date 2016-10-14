package warning

import (
	"github.com/stianeikeland/go-rpio"
)
// RpioImpl wrapper for the rpio package
//
type RpioImpl struct {}

func (r RpioImpl) Open() error {
	return rpio.Open()
}
func (r RpioImpl) Close() error {
	return rpio.Close()
}

func (r RpioImpl) Pin(number int) GpioPin {
	return rpio.Pin(number)
}