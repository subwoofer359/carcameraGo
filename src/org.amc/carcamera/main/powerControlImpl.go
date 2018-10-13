// To signal the PowerBoost card to turn on and stay on
// To detect when external power is off and alert user
package main

import (
	"log"
	"time"

	"org.amc/carcamera/warning"
)

const (
	uSBPOWERON int = 9 // Input pin to receive USB power signal
)

var wAITTIME = 10 * time.Second

type PowerControlImpl struct {
	started    bool
	gpio       warning.Gpio
	usbPowerOn warning.GpioPin
	poweroff   chan bool
}

func (p *PowerControlImpl) SetGPIO(gpio warning.Gpio) {
	p.gpio = gpio
}

func (p *PowerControlImpl) Init() error {
	if err := p.gpio.Open(); err != nil {
		return err
	}
	p.usbPowerOn = p.gpio.Pin(uSBPOWERON)
	p.usbPowerOn.Input()

	p.poweroff = make(chan bool)

	return nil
}

func (p *PowerControlImpl) Start() {
	p.started = true
	go func() {
		for {
			time.Sleep(wAITTIME)
			if p.usbPowerOn.Read() == warning.High {
				log.Println("Power is on")
			} else {
				log.Println("Power is off")
				p.poweroff <- true
				break
			}
		}
	}()
}

func (p PowerControlImpl) IsStarted() bool {
	return p.started
}

func (p *PowerControlImpl) PowerOff() chan bool {
	return p.poweroff
}