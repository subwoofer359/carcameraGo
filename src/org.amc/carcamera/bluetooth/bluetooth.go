package bluetooth

import (
	"github.com/paypal/gatt"
	"github.com/paypal/gatt/linux/cmd"
	"log"
	"fmt"
	"runtime"
)
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
				d.AddService(NewGapService("DashCam"))
				d.AddService(NewGattService())
			
				s1 := NewDashCamService()
			
				d.AddService(s1)
			
				d.AdvertiseNameAndServices("Dash Cam", []gatt.UUID{s1.UUID()})
				
				d.AdvertiseIBeacon(gatt.MustParseUUID("AA6062F098CA42118EC4193EB73CCEB6"), 1, 2, -59)
			default:
		}
	}
	
	d.Init(onStateChanged)
	
	for {
		GetDashCamBTService().Update()
		runtime.Gosched()
	}
}