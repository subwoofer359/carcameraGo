package warning

var open bool

const pinArraySize = 48

type MockGpio struct {
	pins [pinArraySize]GpioPin
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

func (m *MockGpio) AddPin(i int, pin GpioPin) {
	if i < pinArraySize {
		m.pins[i] = pin
	}
}

type MockGpioPin struct {
	State State
	Mode  Mode
}

func (m *MockGpioPin) High() {
	m.State = High
}

func (m *MockGpioPin) Low() {
	m.State = Low
}

func (m *MockGpioPin) Input() {
	m.Mode = Input
}

func (m *MockGpioPin) Output() {
	m.Mode = Output
}

func (m MockGpioPin) Read() State {
	return m.State
}

func NewMockGPIO() *MockGpio {
	return new(MockGpio)
}
