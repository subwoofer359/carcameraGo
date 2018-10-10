package runner

import (
	"os/exec"

	"org.amc/carcamera/storageManager"
)

type CameraCommandFactory interface {
	NewCameraCommand(command string, args []string, storageManager storageManager.StorageManager, exec func(string, ...string) *exec.Cmd) CameraCommand
}
