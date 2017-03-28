package userupdate

import (
	"org.amc/carcamera/bluetooth"
)

type BTService struct {
	 dashService interface { 
	 	SendStatus(val bool)
	 	SendError(errorMsg string)
	 }
}

func (bt *BTService) Init() error {
	bt.dashService = bluetooth.GetDashCamBTService()
	go bluetooth.StartBLE()
	return nil
}

func (bt BTService)	Error(message string) {
	bt.dashService.SendError(message)
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