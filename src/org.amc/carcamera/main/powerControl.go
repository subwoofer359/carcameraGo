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

type PowerControl struct {
	gpio       warning.Gpio
	usbPowerOn warning.GpioPin
	poweroff   chan bool
}

func SetGPIO(p *PowerControl, gpio warning.Gpio) {
	p.gpio = gpio
}

func (p *PowerControl) Init() error {
	if err := p.gpio.Open(); err != nil {
		return err
	}
	p.usbPowerOn = p.gpio.Pin(uSBPOWERON)
	p.usbPowerOn.Input()

	p.poweroff = make(chan bool)

	return nil
}

func (p *PowerControl) Start() {

	for {
		time.Sleep(wAITTIME)
		if p.usbPowerOn.Read() == warning.High {
			log.Println("Power is on")
			p.poweroff <- true
			break
		} else {
			log.Println("Power is off")
		}
	}
}

func (p *PowerControl) Stop() {

}
