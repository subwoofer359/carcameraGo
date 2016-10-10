package main

import (
	"org.amc/carcamera/storageManager"
)

var WebCamApp CameraCommand;

func init() {
	WebCamApp = CameraCommand {
		command: "/usr/bin/raspivid",
		args: []string{},
		storageManager: storageManager.New(),
	}
	
	WebCamApp.storageManager.WorkDir = "/tmp"
	WebCamApp.storageManager.Init()
}


func main() {
}
