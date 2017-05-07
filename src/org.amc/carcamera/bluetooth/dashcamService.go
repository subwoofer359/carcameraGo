package bluetooth

import (
	"log"

	"github.com/paypal/gatt"
)

//=========== DashCam Service ===================

type dashCamBTService struct {
	statusChanged bool
	status        bool
	dsStatus      chan bool
	errorChanged  bool
	dsErrorMsg    chan string
	errorMsg      string
}

var (
	attrDCSUUID = gatt.MustParseUUID("0104988e-59eb-4a00-b051-3e7b2565f631")
	dcBTServ    = new(dashCamBTService)
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

func (d *dashCamBTService) Update() {
	select {
	case status := <-d.dsStatus:
		d.setStatus(status)
	default:
	}

	select {
	case errorMsg := <-d.dsErrorMsg:
		log.Println("Updating Error message")
		d.setErrorMsg(errorMsg)
	default:
	}
}

//GetDashCamBTService returns a struct
//for retrieving values from the Bluetooth Service
func GetDashCamBTService() *dashCamBTService {
	return dcBTServ
}
