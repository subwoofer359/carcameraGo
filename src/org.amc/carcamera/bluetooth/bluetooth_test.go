package bluetooth

import (
	"testing"
	"time"
)

func TestDashCamService(t *testing.T) {
	d := NewDashCamService()
	t.Logf("UUID: %s", d.UUID())
	
	if !d.UUID().Equal(attrDCSUUID) {
		t.Error("DashCam Serice UUID not set")
	}
	
	var check bool = false
	
	for _,char := range d.Characteristics() {
		if char.UUID().Equal(attrSTATUSUUID) {
			check = true
		}
	}
	
	if check == false {
		t.Error("Characteristic Read UUID not set")
	}
}


func TestStartBLE(t *testing.T) {
	go StartBLE()
	time.Sleep(10 * time.Second)
	dcBTServ.SendStatus(true)
	time.Sleep(1 * time.Second)
	dcBTServ.SendStatus(true)
	time.Sleep(1 * time.Second)
	dcBTServ.SendStatus(false)
	dcBTServ.SendError("Help me!")
	time.Sleep(1 * time.Second)
	dcBTServ.SendStatus(true)
	time.Sleep(1 * time.Second)
	dcBTServ.SendStatus(false)
	time.Sleep(60 * time.Second)
}
