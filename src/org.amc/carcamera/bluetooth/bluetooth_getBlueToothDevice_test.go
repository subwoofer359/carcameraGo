package bluetooth

import (
	"github.com/stretchr/testify/assert"
	"github.com/paypal/gatt"
    "testing"
	"errors"
	"os"
	"time"
)

func TestMain(m *testing.M) {
	// Reduce retry delay for testing
	oldTime := BT_WAIT
	BT_WAIT = 1 * time.Millisecond
	v := m.Run()
	BT_WAIT = oldTime
	os.Exit(v)
}

func TestGetBluetooth(t *testing.T) {
	getMockDevice := func (options ...gatt.Option) (gatt.Device, error) {
		mockBT := new(mockBlueToothDevice)
		return mockBT, nil
	}
	d, err := getBluetoothDevice(getMockDevice)
	
	assert.NoError(t, err, "Should be no error")
	
	assert.NotNil(t, d, "Should return value not nil")
}

func TestGetBluetoothOnSecondCall(t *testing.T) {
	var called int = 0
	getMockDeviceOnSecondCall := func(options ...gatt.Option) (gatt.Device, error) {
		if called == 0 {
			called++ 
			return nil, errors.New("test: no device")
		} else {
			mockBT := new(mockBlueToothDevice)
			return mockBT, nil
		}
	} 
	
	d, err := getBluetoothDevice(getMockDeviceOnSecondCall)
	
	assert.NoError(t, err, "Should be no error")
	
	assert.NotNil(t, d, "Should return value not nil")
}

func TestGetBluetoothReturnsNil(t *testing.T) {
	getMockDevice := func(options ...gatt.Option) (gatt.Device, error) {
		return nil, errors.New("test: no device")
	} 
	d, err := getBluetoothDevice(getMockDevice)
	
	assert.Error(t, err, "Should be error")
	
	assert.Nil(t, d, "Should return value nil")
}

func TestGetBluetoothRetriesNTimes(t *testing.T) {
	//Should ignore first try
	var tried int = -1
	getMockDevice := func(options ...gatt.Option) (gatt.Device, error) {
		tried ++
		return nil, errors.New("test: no device")
	} 
	
	d, err := getBluetoothDevice(getMockDevice)
	
	assert.Error(t, err, "Should be error")
	
	assert.Nil(t, d, "Should return value nil")
	
	assert.Equal(t, BT_RETRY, tried, "call to get new device wasn't was equal to set RETRY times")
}
