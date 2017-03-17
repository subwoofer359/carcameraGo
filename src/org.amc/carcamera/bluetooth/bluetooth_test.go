package bluetooth

import (
	"testing"
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
	StartBLE()
}