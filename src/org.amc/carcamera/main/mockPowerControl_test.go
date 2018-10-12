package main

import (
	"org.amc/carcamera/warning"
)

type mockPowerControl struct {
	started  bool
	poweroff chan bool
}

func (m *mockPowerControl) SetGPIO(p *PowerControlImpl, gpio warning.Gpio) {}

func (m *mockPowerControl) Init() error {
	m.poweroff = make(chan bool)
	return nil
}
func (m mockPowerControl) IsStarted() bool { return m.started }

func (m *mockPowerControl) PowerOff() chan bool {
	return m.poweroff
}

func (m *mockPowerControl) Start() {}
