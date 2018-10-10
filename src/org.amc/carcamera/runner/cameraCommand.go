package runner

import (
	"os"

	"org.amc/carcamera/storageManager"
)

//CameraCommand encapsulates the options required by the external camera program 'raspivid'
type CameraCommand interface {
	Args() []string
	Command() string
	SetCommand(command string)
	Process() *os.Process
	Run() error
	StorageManager() storageManager.StorageManager
	SetStorageManager(storageManager storageManager.StorageManager)
}
