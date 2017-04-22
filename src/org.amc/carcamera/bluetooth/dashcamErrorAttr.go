package bluetooth

import (
	"github.com/paypal/gatt"
	"log"
	"org.amc/carcamera/util"
	"time"
)

var (
	attrERRORUUID     = gatt.MustParseUUID("bcc13eef-ad4e-45d3-a2b7-a0b4a4d3d296")
	MESSAGE_SIZE  int = 22
)

func addErrorCharacteristic(s *gatt.Service) {
	c := s.AddCharacteristic(attrERRORUUID)

	c.HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			message := dcBTServ.getErrorMsg()
			if len(message) > MESSAGE_SIZE {
				for _, frag := range util.StringChop(message, MESSAGE_SIZE) {
					rsp.Write([]byte(frag))
				}
			} else {
				rsp.Write([]byte(message))
			}
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
		time.Sleep(NOTIFY_DELAY)
	}
	log.Println("Notify method exited")
}

func (d *dashCamBTService) SendError(errorMsg string) {
	log.Printf("Sending error message(%s)", errorMsg)
	select {
	case d.dsErrorMsg <- errorMsg:
	default:
		log.Println("Error channel is full")
	}
}

func (d *dashCamBTService) setErrorMsg(errorMsg string) {
	d.errorMsg = errorMsg
	d.errorChanged = true
}

func (d dashCamBTService) getErrorMsg() string {
	return d.errorMsg
}
