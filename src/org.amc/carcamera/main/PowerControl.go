package main

import "org.amc/carcamera/warning"

type PowerControl interface {
	SetGPIO(p *PowerControlImpl, gpio warning.Gpio)
	Init() error
	IsStarted() bool
	PowerOff() chan bool
	Start()
}
