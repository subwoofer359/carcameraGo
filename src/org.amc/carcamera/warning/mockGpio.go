package warning

var open bool

type MockGpio struct {
	pins [48]GpioPin
}

func (m *MockGpio) Open() error {
	open = true
	return nil
}

func (m *MockGpio) Close() error {
	open = false
	return nil
}

func (m MockGpio) IsOpen() bool {
	return open
}

func (m *MockGpio) Pin(i int) GpioPin {
	if m.pins[i] == nil {
		m.pins[i] = new(MockGpioPin)
	}
	return m.pins[i]
}

type MockGpioPin struct {
	state State
}

func (m *MockGpioPin) High() {
	m.state = High
}

func (m *MockGpioPin) Low() {
	m.state = Low
}

func (m MockGpioPin) Input() {

}

func (m MockGpioPin) Output() {

}

func (m MockGpioPin) Read() State {
	return m.state
}

func NewMockGPIO() *MockGpio {
	return new(MockGpio)
}
