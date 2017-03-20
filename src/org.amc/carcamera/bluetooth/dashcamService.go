package bluetooth

import (
	"github.com/paypal/gatt"
	"strconv"
	"log"
)

//=========== DashCam Service ===================

type dashCamBTService struct {
	status bool
	dsStatus chan bool
}

var ( 
	attrDCSUUID = gatt.MustParseUUID("0104988e-59eb-4a00-b051-3e7b2565f631")
	attrSTATUSUUID = gatt.MustParseUUID("1f5cf887-d886-484c-8140-8ec407d0e9e6")
	dcBTServ = new (dashCamBTService)
	
)

func init() {
	dcBTServ.dsStatus = make(chan bool, 1)
}

func NewDashCamService() *gatt.Service {
	s := gatt.NewService(attrDCSUUID)
	c := s.AddCharacteristic(attrSTATUSUUID)
	
	c.HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			rsp.Write([]byte(strconv.FormatBool(getStatus())))
		})
	c.HandleNotifyFunc(
		func(r gatt.Request, n gatt.Notifier) {
			notify(n)
		})
	c.AddDescriptor(gatt.UUID16(0x2901)).SetValue([]byte("Dashcam status"))
	
	c.AddDescriptor(gatt.UUID16(0x2904)).SetValue([]byte{0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00})
	
	return s
}

func (d *dashCamBTService) SendStatus(status bool) {
	d.dsStatus <- status
}

func GetDashCamBTService() *dashCamBTService {
	return dcBTServ
}

func getStatus() bool {
	select {
			case status := <- dcBTServ.dsStatus:
			dcBTServ.status = status				
			default:
		}
	return dcBTServ.status
}

func notify(n gatt.Notifier) {
	for !n.Done() {
		select {
			case status := <- dcBTServ.dsStatus:
				log.Printf("Notify: written message: %s", strconv.FormatBool(status))
				dcBTServ.status = status
				n.Write([]byte(strconv.FormatBool(status)))
			default:
				//log.Println("Notify:No message received")
		}
	}
	log.Println("Notify method exited")
}