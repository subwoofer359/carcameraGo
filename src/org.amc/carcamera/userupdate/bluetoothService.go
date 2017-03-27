package userupdate

import (
	"org.amc/carcamera/bluetooth"
)

type BTService struct {
	 dashService interface{ SendStatus(val bool)}
}

func (bt *BTService) Init() error {
	go bluetooth.StartBLE()
	return nil
}

func (bt BTService)	Error(message string) {
	//Todo
}

func (bt *BTService) Started() {
	bt.dashService.SendStatus(true)
}

func (bt *BTService) Stopped() {
	bt.dashService.SendStatus(false)
}

func (bt *BTService) Close() {
	//Todo
}