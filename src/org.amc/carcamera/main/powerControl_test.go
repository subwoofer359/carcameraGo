package main

import (
	"log"
	"testing"

	"org.amc/carcamera/warning"
)

func TestPowerControlInit(t *testing.T) {
	m := new(warning.MockGpio)

	tpin := &warning.MockGpioPin{}

	m.AddPin(usbPowerOn, tpin)

	powerControl := PowerControl{
		gpio: m,
	}

	if err := powerControl.Init(); err != nil {
		t.Error(err)
	}

	if !m.IsOpen() {
		t.Error("GPIO system not opened")
	}

	log.Println(tpin.Mode)

	if tpin.Mode != warning.Input {
		t.Error("GPIO pin not set to Input mode")
	}
}
