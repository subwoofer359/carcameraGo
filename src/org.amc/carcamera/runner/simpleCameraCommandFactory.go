package runner

import (
	"os/exec"

	"org.amc/carcamera/storageManager"
)

type SimpleCameraCommandFactory struct {
}

func (c SimpleCameraCommandFactory) NewCameraCommand(command string, args []string, storageManager storageManager.StorageManager, exec func(string, ...string) *exec.Cmd) CameraCommand {
	return &CameraCommandImpl{
		command:        command,
		args:           args,
		storageManager: storageManager,
		exec:           exec,
	}
}
