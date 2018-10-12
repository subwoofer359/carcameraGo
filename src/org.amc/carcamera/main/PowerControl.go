package main

import (
	"errors"

	"org.amc/carcamera/warning"
)

var ErrPowerFault = errors.New("Power Control Interrupt")

type PowerControl interface {
	SetGPIO(p *PowerControlImpl, gpio warning.Gpio)
	Init() error
	IsStarted() bool
	PowerOff() chan bool
	Start()
}
