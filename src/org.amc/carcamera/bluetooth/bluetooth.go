package bluetooth

import (
	"github.com/paypal/gatt"
	"github.com/paypal/gatt/linux/cmd"
	"log"
	"fmt"
)
//=========== GAP ==============
var (
	attrGAPUUID = gatt.UUID16(0x1800)

	attrDeviceNameUUID        = gatt.UUID16(0x2A00)
	attrAppearanceUUID        = gatt.UUID16(0x2A01)
	attrPeripheralPrivacyUUID = gatt.UUID16(0x2A02)
	attrReconnectionAddrUUID  = gatt.UUID16(0x2A03)
	attrPeferredParamsUUID    = gatt.UUID16(0x2A04)
)

var gapCharAppearanceGenericComputer = []byte{0x00, 0x80}

func NewGapService(name string) *gatt.Service {
	s := gatt.NewService(attrGAPUUID)
	s.AddCharacteristic(attrDeviceNameUUID).SetValue([]byte(name))
	s.AddCharacteristic(attrAppearanceUUID).SetValue(gapCharAppearanceGenericComputer)
	s.AddCharacteristic(attrPeripheralPrivacyUUID).SetValue([]byte{0x00})
	s.AddCharacteristic(attrReconnectionAddrUUID).SetValue([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	s.AddCharacteristic(attrPeferredParamsUUID).SetValue([]byte{0x06, 0x00, 0x06, 0x00, 0x00, 0x00, 0xd0, 0x07})
	return s
}

// ================ GATT =================

var (
	attrGATTUUID           = gatt.UUID16(0x1801)
	attrServiceChangedUUID = gatt.UUID16(0x2A05)
)

func NewGattService() *gatt.Service {
	s := gatt.NewService(attrGATTUUID)
	s.AddCharacteristic(attrServiceChangedUUID).HandleNotifyFunc(
		func(r gatt.Request, n gatt.Notifier) {
			go func() {
				log.Printf("TODO: indicate client when the services are changed")
			}()
		})
	return s
}

//=========== DashCam Service ===================

var ( 
	attrDCSUUID = gatt.MustParseUUID("0104988e-59eb-4a00-b051-3e7b2565f631")
	attrSTATUSUUID = gatt.MustParseUUID("1f5cf887-d886-484c-8140-8ec407d0e9e6")
)

func NewDashCamService() *gatt.Service {
	s := gatt.NewService(attrDCSUUID)
	c := s.AddCharacteristic(attrSTATUSUUID)
	
	c.HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			rsp.Write([]byte("true"))
		})
	c.HandleNotifyFunc(
		func(r gatt.Request, n gatt.Notifier) {
			notify(n)	
		})
	c.AddDescriptor(gatt.UUID16(0x2901)).SetValue([]byte("Dashcam status"))
	
	c.AddDescriptor(gatt.UUID16(0x2904)).SetValue([]byte{0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00})
	
	return s
}

func notify(n gatt.Notifier) {
	for !n.Done() {
		n.Write([]byte("Notifying you"))
	}
}


//============== Main Service ==================
var DefaultClientOptions = []gatt.Option{
	gatt.LnxMaxConnections(1),
	gatt.LnxDeviceID(-1, true),
}

var DefaultServerOptions = []gatt.Option{
	gatt.LnxMaxConnections(2),
	gatt.LnxDeviceID(-1, true),
	gatt.LnxSetAdvertisingParameters(&cmd.LESetAdvertisingParameters{
		AdvertisingIntervalMin: 0x00f4,
		AdvertisingIntervalMax: 0x00f4,
		AdvertisingChannelMap:  0x7,
	}),
}

func StartBLE() {
	d, err := gatt.NewDevice(DefaultServerOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s", err)
	}
	
	d.Handle(
		gatt.CentralConnected(func(c gatt.Central) { fmt.Println("Connect: ", c.ID()) }),
		gatt.CentralDisconnected(func(c gatt.Central) { fmt.Println("Disconnect: ", c.ID()) }),
	)
	
	onStateChanged := func(d gatt.Device, s gatt.State) {
		fmt.Printf("State: %s\n", s)
	
		switch s {
			case gatt.StatePoweredOn:
				d.AddService(NewGapService("Gopher"))
				d.AddService(NewGattService())
			
				s1 := NewDashCamService()
			
				d.AddService(s1)
			
				d.AdvertiseNameAndServices("DashCam", []gatt.UUID{s1.UUID()})
				
				d.AdvertiseIBeacon(gatt.MustParseUUID("AA6062F098CA42118EC4193EB73CCEB6"), 1, 2, -59)
			default:
		}
	}
	
	d.Init(onStateChanged)
	select {}
}