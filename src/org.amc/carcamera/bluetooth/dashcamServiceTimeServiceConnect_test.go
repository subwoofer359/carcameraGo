package bluetooth

import (
	"testing"

	"github.com/paypal/gatt"
	"github.com/stretchr/testify/assert"
)

func TestScanForTimeService(t *testing.T) {
	mockDevice := new(mockBlueToothDevice)
	startScanning(mockDevice)
	assert.True(t, mockDevice.scanning)
}

func TestTimeServiceDiscoveredHandler(t *testing.T) {
	var advertisement = new(gatt.Advertisement)
	var testName = "MotoG4"

	var aTimeDevice = new(myPeripheral)
	aTimeDevice.name = testName

	timeServiceDiscoveredHandler(aTimeDevice, advertisement, 0)
}

type myPeripheral struct {
	name string
}

// Device returns the underlying device.
func (p myPeripheral) Device() gatt.Device {
	return nil
}

// ID is the platform specific unique ID of the remote peripheral, e.g. MAC for Linux, gatt.Peripheral UUID for MacOS.
func (p myPeripheral) ID() string {
	return "1"
}

// Name returns the name of the remote peripheral.
// This can be the advertised name, if exists, or the GAP device name, which takes priority
func (p myPeripheral) Name() string {
	return p.name
}

// Services returnns the services of the remote peripheral which has been discovered.
func (p myPeripheral) Services() []*gatt.Service {
	return []*gatt.Service{}
}

// DiscoverServices discover the specified services of the remote peripheral.
// If the specified services is set to nil, all the available services of the remote peripheral are returned.
func (p myPeripheral) DiscoverServices(s []gatt.UUID) ([]*gatt.Service, error) {
	return []*gatt.Service{}, nil
}

// DiscoverIncludedServices discovers the specified included services of a service.
// If the specified services is set to nil, all the included services of the service are returned.
func (p myPeripheral) DiscoverIncludedServices(ss []gatt.UUID, s *gatt.Service) ([]*gatt.Service, error) {
	return []*gatt.Service{}, nil
}

// DiscoverCharacteristics discovers the specified characteristics of a service.
// If the specified characterstics is set to nil, all the characteristic of the service are returned.
func (p myPeripheral) DiscoverCharacteristics(c []gatt.UUID, s *gatt.Service) ([]*gatt.Characteristic, error) {
	return []*gatt.Characteristic{}, nil
}

// DiscoverDescriptors discovers the descriptors of a characteristic.
// If the specified descriptors is set to nil, all the descriptors of the characteristic are returned.
func (p myPeripheral) DiscoverDescriptors(d []gatt.UUID, c *gatt.Characteristic) ([]*gatt.Descriptor, error) {
	return []*gatt.Descriptor{}, nil
}

// ReadCharacteristic retrieves the value of a specified characteristic.
func (p myPeripheral) ReadCharacteristic(c *gatt.Characteristic) ([]byte, error) {
	return []byte{}, nil
}

// ReadLongCharacteristic retrieves the value of a specified characteristic that is longer than the
// MTU.
func (p myPeripheral) ReadLongCharacteristic(c *gatt.Characteristic) ([]byte, error) {
	return []byte{}, nil
}

// ReadDescriptor retrieves the value of a specified characteristic descriptor.
func (p myPeripheral) ReadDescriptor(d *gatt.Descriptor) ([]byte, error) {
	return []byte{}, nil
}

// WriteCharacteristic writes the value of a characteristic.
func (p myPeripheral) WriteCharacteristic(c *gatt.Characteristic, b []byte, noRsp bool) error {
	return nil
}

// WriteDescriptor writes the value of a characteristic descriptor.
func (p myPeripheral) WriteDescriptor(d *gatt.Descriptor, b []byte) error {
	return nil
}

// SetNotifyValue sets notifications for the value of a specified characteristic.
func (p myPeripheral) SetNotifyValue(c *gatt.Characteristic, f func(*gatt.Characteristic, []byte, error)) error {
	return nil
}

// SetIndicateValue sets indications for the value of a specified characteristic.
func (p myPeripheral) SetIndicateValue(c *gatt.Characteristic, f func(*gatt.Characteristic, []byte, error)) error {
	return nil
}

// ReadRSSI retrieves the current RSSI value for the remote peripheral.
func (p myPeripheral) ReadRSSI() int {
	return 0
}

// SetMTU sets the mtu for the remote peripheral.
func (p myPeripheral) SetMTU(mtu uint16) error {
	return nil
}
