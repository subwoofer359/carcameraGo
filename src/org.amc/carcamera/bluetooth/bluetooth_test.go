package bluetooth

import (
	"github.com/stretchr/testify/assert"
	C "org.amc/carcamera/constants"
	"testing"
	//"time"
)

func TestDashCamService(t *testing.T) {
	d := NewDashCamService()
	t.Logf("UUID: %s", d.UUID())

	if !d.UUID().Equal(attrDCSUUID) {
		t.Error("DashCam Serice UUID not set")
	}

	var check bool = false

	for _, char := range d.Characteristics() {
		if char.UUID().Equal(attrSTATUSUUID) {
			check = true
		}
	}

	if check == false {
		t.Error("Characteristic Read UUID not set")
	}
}

func TestSetServiceNames(t *testing.T) {
	var context = map[string]interface{}{}
	gap := "Testing GAP Service"
	gatt := "Testing GATT Service"
	context[C.GAP_SERVICE_NAME] = gap

	assert.Equal(t, GAP_SERVICE_NAME, "DashCam")
	setServiceNames(context)
	assert.Equal(t, GAP_SERVICE_NAME, gap)

	assert.Equal(t, GATT_SERVICE_NAME, "Dash Cam")
	context[C.GATT_SERVICE_NAME] = gatt
	setServiceNames(context)
	assert.Equal(t, GATT_SERVICE_NAME, gatt)
}

//func TestStartBLE(t *testing.T) {
//	go StartBLE()
//	time.Sleep(10 * time.Second)
//	dcBTServ.SendStatus(true)
//	time.Sleep(1 * time.Second)
//	dcBTServ.SendStatus(true)
//	time.Sleep(1 * time.Second)
//	dcBTServ.SendStatus(false)
//	dcBTServ.SendError("Help me!")
//	time.Sleep(1 * time.Second)
//	dcBTServ.SendStatus(true)
//	time.Sleep(1 * time.Second)
//	dcBTServ.SendStatus(false)
//	time.Sleep(60 * time.Second)
//}
