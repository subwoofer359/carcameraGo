// To signal the PowerBoost card to turn on and stay on
// To detect when external power is off and alert user
package main

import (
	"log"
	"time"

	"org.amc/carcamera/warning"
)

const (
	usbPowerOn int = 9 // Input pin to receive USB power signal
	wAITTIME       = 10 * time.Second
)

type PowerControl struct {
	gpio       warning.Gpio
	usbPowerOn warning.GpioPin
}

func SetGPIO(p *PowerControl, gpio warning.Gpio) {
	p.gpio = gpio
}

func (p *PowerControl) Init() error {
	if err := p.gpio.Open(); err != nil {
		return err
	}
	p.usbPowerOn = p.gpio.Pin(usbPowerOn)
	p.usbPowerOn.Input()
	return nil
}

func (p *PowerControl) Start() {

	for {
		time.Sleep(wAITTIME)
		if p.usbPowerOn.Read() == warning.High {
			log.Println("Power is on")
		} else {
			log.Println("Power is off")
		}
	}
}

func (p *PowerControl) Stop() {

}
