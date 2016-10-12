package main

import (
	"org.amc/carcamera/storageManager"
	"os/exec"
)

var WebCamApp CameraCommandImpl;
var runner Runner 

func init() {
	WebCamApp = CameraCommandImpl {
		command: "/usr/bin/raspivid",
		args: []string{},
		storageManager: storageManager.New(),
		exec: exec.Command,
	}
	
	WebCamApp.storageManager.SetWorkDir("/tmp")
	WebCamApp.storageManager.Init()
}


func main() {

}
