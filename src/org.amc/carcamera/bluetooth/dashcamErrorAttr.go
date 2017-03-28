package bluetooth

import (
	"github.com/paypal/gatt"
	"log"
)

var (
	attrERRORUUID = gatt.MustParseUUID("bcc13eef-ad4e-45d3-a2b7-a0b4a4d3d296")
)

func addErrorCharacteristic(s *gatt.Service) {
	c := s.AddCharacteristic(attrERRORUUID)
	
	c.HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			rsp.Write([]byte(dcBTServ.getErrorMsg()))
		})
	c.HandleNotifyFunc(
		func(r gatt.Request, n gatt.Notifier) {
			notifyError(n, dcBTServ)
		})
	
	c.AddDescriptor(gatt.UUID16(0x2901)).SetValue([]byte("Error message"))
	
	c.AddDescriptor(gatt.UUID16(0x2904)).SetValue([]byte{0x1A, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00})
	
}

func notifyError(n gatt.Notifier, d *dashCamBTService) {
	for !n.Done() {
		if d.errorMsg != "" && d.errorChanged {
			log.Printf("Notify: written message: %s", d.errorMsg)
			n.Write([]byte(d.errorMsg))
			d.errorChanged = false
		}
	}
	log.Println("Notify method exited")
}

func (d *dashCamBTService) SendError(errorMsg string) {
	log.Printf("Sending error message(%s)", errorMsg)
	d.dsErrorMsg <- errorMsg
}

func (d *dashCamBTService) setErrorMsg(errorMsg string) {
	d.errorMsg = errorMsg
	d.errorChanged = true
}

func (d dashCamBTService) getErrorMsg() string {
	return d.errorMsg
}
