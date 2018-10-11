package main

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"org.amc/carcamera/warning"
)

var powerControl PowerControlImpl
var m *warning.MockGpio
var tpin *warning.MockGpioPin

func powerControlSetup() {

	m = new(warning.MockGpio)

	tpin = &warning.MockGpioPin{
		Mode: warning.Input,
	}

	m.AddPin(uSBPOWERON, tpin)

	powerControl = PowerControlImpl{
		gpio: m,
	}
}

func TestPowerControlInit(t *testing.T) {

	powerControlSetup()

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

	//Has poweroff been initialised

	if powerControl.poweroff == nil {
		t.Error("Poweroff channel not initialised")
	}

	if powerControl.usbPowerOn == nil {
		t.Error("UsbPowerOn shouldn't be nil")
	}
}

func TestPowerControlStart(t *testing.T) {
	powerControlSetup()

	wAITTIME = 10 * time.Millisecond

	if err := powerControl.Init(); err != nil {
		t.Error(err)
	}

	assert.Equal(t, warning.Low, powerControl.usbPowerOn.Read())

	powerControl.Start()

	time.Sleep(1 * time.Second)
	//Set pin high
	powerControl.usbPowerOn = &warning.MockGpioPin{State: warning.High}

	select {
	case state := <-powerControl.poweroff:
		log.Println(state)
	case <-time.After(3 * time.Second):
		log.Fatal("Test timed out")
	}
}
