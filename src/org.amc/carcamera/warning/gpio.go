package warning

import (
	"github.com/stianeikeland/go-rpio"
)
type Gpio interface {
	Open() error
	Close()
	Pin(int) GpioPin
}

type GpioPin interface {
	High()
	Low()
	Output()
	Read() rpio.State
}
