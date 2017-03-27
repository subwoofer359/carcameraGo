package bluetooth

import (
	"github.com/paypal/gatt"
	"strconv"
	"log"
)

//=========== DashCam Service ===================

type dashCamBTService struct {
	statusChanged bool
	status bool
	dsStatus chan bool
	errorChanged bool
	dsErrorMsg chan string
	errorMsg string
}

var ( 
	attrDCSUUID = gatt.MustParseUUID("0104988e-59eb-4a00-b051-3e7b2565f631")
	attrSTATUSUUID = gatt.MustParseUUID("1f5cf887-d886-484c-8140-8ec407d0e9e6")
	attrERRORUUID = gatt.MustParseUUID("bcc13eef-ad4e-45d3-a2b7-a0b4a4d3d296")
	dcBTServ = new (dashCamBTService)
	
)

func init() {
	dcBTServ.dsStatus = make(chan bool, 1)
	dcBTServ.dsErrorMsg = make(chan string, 1)
}

func NewDashCamService() *gatt.Service {
	s := gatt.NewService(attrDCSUUID)
	
	addStatusCharacteristic(s)
	
	addErrorCharacteristic(s)
	return s
}

/*
 * Set up Characteristic for Status
 */

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
	}
	log.Println("Notify method exited")
}

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

// ============= Status Characteristic =============

func (d *dashCamBTService) SendStatus(status bool) {
	log.Println("Sending Status")
	d.dsStatus <- status
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
// ============= Error Characteristic =============

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

func (d *dashCamBTService) Update() {
	select {
		case status := <- d.dsStatus:
			d.setStatus(status)						
		default:
	}
	
	select {
		case errorMsg := <- d.dsErrorMsg:
			log.Println("Updating Error message")
			d.setErrorMsg(errorMsg)
		default:
	}
}

func GetDashCamBTService() *dashCamBTService {
	return dcBTServ
}