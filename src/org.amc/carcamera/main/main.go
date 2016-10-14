package main

import (
	"org.amc/carcamera/storageManager"
	"org.amc/carcamera/warning"
	"os/exec"
	"time"
	"log"
)


const appTimeout time.Duration = time.Duration(7) * time.Minute

var (

	runner Runner

	lights warning.UserDisplay
	
	WebCamApp *CameraCommandImpl
)

func init() {
	log.Println("Starting WebCam Program")
	lights.SetGPIO(warning.RpioImpl{})
}

func main() {
	lights.Open()
	defer lights.Close()
	
	WebCamApp = createWebCamCommand() 
	
	initStorageManager()

	lights.Ok()
	
	start()
}

func createWebCamCommand() *CameraCommandImpl {
	return &CameraCommandImpl {
		command: "/usr/bin/raspivid",
		args: []string{},
		storageManager: storageManager.New(),
		exec: exec.Command,
	}
}

func initStorageManager() {
	WebCamApp.storageManager.SetWorkDir("/tmp")
	if err := WebCamApp.storageManager.Init(); err != nil {
		lights.Error()
		log.Fatal(err)
	}
}

func start() {
	for {
		runner := New(appTimeout)
		runner.add(WebCamApp)
		
		err := runner.Start()
		
		if err != nil && err.Error() != "completed" {
			lights.Error()
			log.Fatal(err)
		}
	}
}