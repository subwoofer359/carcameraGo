package powercontrol

import (
	"log"
	"time"

	"org.amc/carcamera/warning"
)

//mockPowerControl PowerControl object that doesnt nothing
type MockPowerControl struct {
	started  bool
	poweroff chan bool
}

func (m *MockPowerControl) SetGPIO(gpio warning.Gpio) {}

func (m *MockPowerControl) Init() error {
	m.poweroff = make(chan bool)
	return nil
}
func (m MockPowerControl) IsStarted() bool { return m.started }

func (m *MockPowerControl) PowerOff() chan bool {
	return m.poweroff
}

func (m *MockPowerControl) Start() {}

//mockPowerControl PowerControl object that sends a poweroff message after a period time
type PPowerControl struct {
	MockPowerControl
}

//Start starts a thread and sends a message after a set time
func (p *PPowerControl) Start() {
	log.Println("Start called on PowerControl")
	go func() {
		log.Println("PowerControl poweroff message sent")
		time.Sleep(10 * time.Millisecond)
		p.poweroff <- true
	}()
}
