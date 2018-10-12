package main

import (
	"log"
	"time"

	"org.amc/carcamera/warning"
)

//mockPowerControl PowerControl object that doesnt nothing
type mockPowerControl struct {
	started  bool
	poweroff chan bool
}

func (m *mockPowerControl) SetGPIO(gpio warning.Gpio) {}

func (m *mockPowerControl) Init() error {
	m.poweroff = make(chan bool)
	return nil
}
func (m mockPowerControl) IsStarted() bool { return m.started }

func (m *mockPowerControl) PowerOff() chan bool {
	return m.poweroff
}

func (m *mockPowerControl) Start() {}

//mockPowerControl PowerControl object that sends a poweroff message after a period time
type pPowerControl struct {
	mockPowerControl
}

//Start starts a thread and sends a message after a set time
func (p *pPowerControl) Start() {
	log.Println("Start called on PowerControl")
	go func() {
		log.Println("PowerControl poweroff message sent")
		time.Sleep(10 * time.Millisecond)
		p.poweroff <- true
	}()
}
