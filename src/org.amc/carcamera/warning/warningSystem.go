package warning

const (
	GreenLED int = 17

	RedLED int = 22
)

type UserDisplay struct {
	gpio Gpio
}

func (u *UserDisplay) SetGPIO(gpio Gpio) {
	u.gpio = gpio
}

func (u *UserDisplay) Ok() {
	if !isHigh(u, RedLED) {
		setHigh(u, GreenLED)
	}
}

// PowerError On error turns green light off
//To replace with a method to make Green Led flash
func (u *UserDisplay) PowerError() {
	setLow(u, GreenLED)
}

func (u *UserDisplay) Error() {
	setHigh(u, RedLED)
}

func (u *UserDisplay) Reset() {
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
}

func (u *UserDisplay) Close() {
	setLow(u, GreenLED)
	u.gpio.Close()
}

func setHigh(u *UserDisplay, colour int) {
	u.gpio.Pin(colour).High()
}

func setLow(u *UserDisplay, colour int) {
	u.gpio.Pin(colour).Low()
}

func read(u *UserDisplay, colour int) State {
	return u.gpio.Pin(colour).Read()
}

func isHigh(u *UserDisplay, colour int) bool {
	return read(u, colour) == High
}
