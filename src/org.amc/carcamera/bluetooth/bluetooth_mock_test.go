package bluetooth

import "github.com/paypal/gatt"

type mockBlueToothDevice struct {}

func (m *mockBlueToothDevice) Init(stateChanged func(gatt.Device, gatt.State)) error {
	return nil
}

    
func (m *mockBlueToothDevice) Advertise(a *gatt.AdvPacket) error {
	return nil
}

func (m *mockBlueToothDevice) AdvertiseNameAndServices(name string, ss []gatt.UUID) error {
	return nil
}

func (m *mockBlueToothDevice) AdvertiseIBeaconData(b []byte) error {
	return nil
}

func (m *mockBlueToothDevice) AdvertiseIBeacon(u gatt.UUID, major, minor uint16, pwr int8) error {
	return nil
}

func (m *mockBlueToothDevice) StopAdvertising() error {
	return nil
}
func (m *mockBlueToothDevice) RemoveAllServices() error {
	return nil
}

func (m *mockBlueToothDevice) AddService(s *gatt.Service) error {
	return nil
}

func (m *mockBlueToothDevice) SetServices(ss []*gatt.Service) error {
	return nil
}
func (m *mockBlueToothDevice)  Scan(ss []gatt.UUID, dup bool) {
	
}

func (m *mockBlueToothDevice) StopScanning() {}

func (m *mockBlueToothDevice) Connect(p gatt.Peripheral) {}

func (m *mockBlueToothDevice) CancelConnection(p gatt.Peripheral) {}

func (m *mockBlueToothDevice) Handle(h ...gatt.Handler) {}

func (m *mockBlueToothDevice) Option(o ...gatt.Option) error {
	return nil
}
