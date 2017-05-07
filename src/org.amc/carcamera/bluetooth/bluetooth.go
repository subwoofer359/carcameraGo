package bluetooth

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/paypal/gatt"
	"github.com/paypal/gatt/linux/cmd"
	C "org.amc/carcamera/constants"
)

//============== Main Service ==================

var (
	DefaultClientOptions = []gatt.Option{
		gatt.LnxMaxConnections(1),
		gatt.LnxDeviceID(-1, true),
	}

	DefaultServerOptions = []gatt.Option{
		gatt.LnxMaxConnections(2),
		gatt.LnxDeviceID(-1, true),
		gatt.LnxSetAdvertisingParameters(&cmd.LESetAdvertisingParameters{
			AdvertisingIntervalMin: 0x00f4,
			AdvertisingIntervalMax: 0x00f4,
			AdvertisingChannelMap:  0x7,
		}),
	}

	UPDATE_DELAY = 10 * time.Second

	NOTIFY_DELAY = 2 * time.Second

	GATT_SERVICE_NAME = "Dash Cam"

	GAP_SERVICE_NAME = "DashCam"

	BT_WAIT = 1 * time.Second

	BT_RETRY = 5

	BT_DEVICE_ERROR = errors.New("Could not open device")

	currentTimeServiceUUID = gatt.UUID16(0x1805)
)

func getBluetoothDevice(newDevice func(opts ...gatt.Option) (gatt.Device, error)) (gatt.Device, error) {
	var (
		d     gatt.Device
		btErr error
	)
	d, btErr = newDevice(DefaultServerOptions...)
	if btErr != nil {
		log.Printf("Failed to open device, err: %s", btErr)
		for i := 0; i < BT_RETRY; i++ {
			time.Sleep(BT_WAIT)
			d, btErr = newDevice(DefaultServerOptions...)
			if btErr != nil {
				log.Printf("Failed to open device, err: %s", btErr)
			} else {
				break
			}
		}
	}

	if d == nil {
		log.Println("Could not open bluetooth device")
		return d, BT_DEVICE_ERROR
	}

	return d, nil

}

// StartBLE starts Bluetooth Services
func StartBLE(context map[string]interface{}) {
	d, err := getBluetoothDevice(gatt.NewDevice)

	if err != nil {
		log.Printf("%s", err.Error())
		return
	}

	setServiceNames(context)

	d.Handle(
		gatt.CentralConnected(func(c gatt.Central) { fmt.Println("Connect: ", c.ID()) }),
		gatt.CentralDisconnected(func(c gatt.Central) { fmt.Println("Disconnect: ", c.ID()) }),
		gatt.PeripheralDiscovered(timeServiceDiscoveredHandler),
	)

	onStateChanged := func(d gatt.Device, s gatt.State) {
		fmt.Printf("State: %s\n", s)

		switch s {
		case gatt.StatePoweredOn:
			d.AddService(NewGapService(GAP_SERVICE_NAME))
			d.AddService(NewGattService())

			s1 := NewDashCamService()

			d.AddService(s1)

			d.AdvertiseNameAndServices(GATT_SERVICE_NAME, []gatt.UUID{s1.UUID()})

			d.AdvertiseIBeacon(gatt.MustParseUUID("AA6062F098CA42118EC4193EB73CCEB6"), 1, 2, -59)
		default:
		}
	}

	d.Init(onStateChanged)

	for {
		GetDashCamBTService().Update()
		runtime.Gosched()
		time.Sleep(UPDATE_DELAY)
	}
}

func setServiceNames(context map[string]interface{}) {
	if context != nil {
		if context[C.GAP_SERVICE_NAME] != nil {
			GAP_SERVICE_NAME = context[C.GAP_SERVICE_NAME].(string)
		}

		if context[C.GATT_SERVICE_NAME] != nil {
			GATT_SERVICE_NAME = context[C.GATT_SERVICE_NAME].(string)
		}
	}
}

func startScanning(device gatt.Device) {
	device.Scan([]gatt.UUID{currentTimeServiceUUID}, false)
}

func timeServiceDiscoveredHandler(peripheral gatt.Peripheral, advertisement *gatt.Advertisement, c int) {
	log.Println("Peripheral:" + peripheral.Name())
}
