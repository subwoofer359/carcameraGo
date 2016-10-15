package warning

import "github.com/stianeikeland/go-rpio"

var open bool

type mockGpio struct {
	pins [48]GpioPin
}

func (m *mockGpio) Open() error {
	open = true
	return nil
}

func (m *mockGpio) Close() error {
	open = false
	return nil
}

func (m mockGpio) IsOpen() bool {
	return open
}


func (m *mockGpio) Pin(i int) GpioPin {
	if m.pins[i] == nil {
		m.pins[i] = new(mockGpioPin)
	}
	return m.pins[i]
}

type mockGpioPin struct {
	state rpio.State
}

func (m *mockGpioPin) High() {
	m.state = rpio.High
}

func (m *mockGpioPin) Low() {
	m.state = rpio.Low
}

func (m mockGpioPin) Output() {
	
}

func (m mockGpioPin) Read() rpio.State {
	return m.state
}



func NewMockGPIO() *mockGpio {
	return new(mockGpio)
}


