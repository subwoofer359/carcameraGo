package warning

//State of the GPIO pin
type State uint8

const (
	//High pin is in the high state
	High State = 1
	//Low pin is in the low state
	Low State = 0
)

//Gpio interface to the GPIO system
type Gpio interface {
	Open() error
	Close() error
	Pin(int) GpioPin
}

//GpioPin interface of a GPIO pin
type GpioPin interface {
	High()
	Low()
	Input()
	Output()
	Read() State
}
