package userupdate

import (
	 "org.amc/carcamera/warning"
)

type LEDService struct {
	led warning.UserDisplay
}

func (l *LEDService) Init() error {
	l.led.Open()
	l.led.Reset()
	return nil
}

func (l LEDService)	Error(message string) {
	l.led.Error()
}

func (l *LEDService) Started() {
	l.led.Ok()
}

func (l *LEDService) Stopped() {
	l.led.Close()
}

func (l *LEDService) SetGPIO(gpio warning.Gpio) {
	l.led.SetGPIO(gpio)
}


