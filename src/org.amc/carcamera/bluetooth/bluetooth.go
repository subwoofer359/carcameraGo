package bluetooth

import (
	"errors"
	"fmt"
	"github.com/paypal/gatt"
	"github.com/paypal/gatt/linux/cmd"
	"log"
	C "org.amc/carcamera/constants"
	"runtime"
	"time"
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

	GATT_SERVICE_NAME string = "Dash Cam"

	GAP_SERVICE_NAME string = "DashCam"

	BT_WAIT = 1 * time.Second

	BT_RETRY = 5

	BT_DEVICE_ERROR = errors.New("Could not open device")
)

func getBluetoothDevice(newDevice func(opts ...gatt.Option) (gatt.Device, error)) (gatt.Device, error) {
	var (
		d      gatt.Device
		bt_err error
	)
	d, bt_err = newDevice(DefaultServerOptions...)
	if bt_err != nil {
		log.Printf("Failed to open device, err: %s", bt_err)
		for i := 0; i < BT_RETRY; i++ {
			time.Sleep(BT_WAIT)
			d, bt_err = newDevice(DefaultServerOptions...)
			if bt_err != nil {
				log.Printf("Failed to open device, err: %s", bt_err)
			} else {
				break
			}
		}
	}

	if d == nil {
		log.Println("Could not open bluetooth device")
		return d, BT_DEVICE_ERROR
	} else {
		return d, nil
	}
}

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
