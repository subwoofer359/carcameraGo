package userupdate

import (
    "testing"
    "org.amc/carcamera/bluetooth"
)

func TestBTServiceStarted(t *testing.T) {
	service := new (BTService)
	service.dashService = bluetooth.GetDashCamBTService()
	service.Started()
}

func TestBTServiceInitialised(t *testing.T) {
	service := new (BTService)
	service.dashService = bluetooth.GetDashCamBTService()
	if err := service.Init(); err != nil {
		t.Error("Error when initialising bluetooth")
	}
	service.Started()
	service.Started()
	service.Started()
}


