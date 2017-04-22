package bluetooth

import (
	"github.com/paypal/gatt"
	"log"
	"strconv"
	"time"
)

var (
	attrSTATUSUUID = gatt.MustParseUUID("1f5cf887-d886-484c-8140-8ec407d0e9e6")
)

func addStatusCharacteristic(s *gatt.Service) {
	c := s.AddCharacteristic(attrSTATUSUUID)

	c.HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			rsp.Write([]byte(strconv.FormatBool(dcBTServ.getStatus())))
		})
	c.HandleNotifyFunc(
		func(r gatt.Request, n gatt.Notifier) {
			notifyStatus(n, dcBTServ)
		})
	c.AddDescriptor(gatt.UUID16(0x2901)).SetValue([]byte("Dashcam status"))

	c.AddDescriptor(gatt.UUID16(0x2904)).SetValue([]byte{0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00})
}

func notifyStatus(n gatt.Notifier, d *dashCamBTService) {
	for !n.Done() {
		if d.statusChanged == true {
			var statusStr string = strconv.FormatBool(d.getStatus())
			log.Printf("Notify: written message: %s", statusStr)
			n.Write([]byte(statusStr))
			d.statusChanged = false

		}
		time.Sleep(NOTIFY_DELAY)
	}
	log.Println("Notify method exited")
}

// ============= Status Characteristic =============

func (d *dashCamBTService) SendStatus(status bool) {
	log.Println("Sending Status")
	select {
	case d.dsStatus <- status:
	default:
		log.Println("Status channel is full")
	}
}

func (d *dashCamBTService) getStatus() bool {
	log.Println("Getting status")

	return d.status
}

func (d *dashCamBTService) setStatus(status bool) {
	if d.status != status {
		d.statusChanged = true
	}
	d.status = status
}
