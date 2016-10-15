package main

import (
	"org.amc/carcamera/storageManager"
	"org.amc/carcamera/warning"
	"os/exec"
	"time"
	"log"
)


const appTimeout time.Duration = time.Duration(7) * time.Minute

type app struct {
	runner *Runner
	lights warning.UserDisplay
	WebCamApp *CameraCommandImpl
}

func (a *app) Init() {
	log.Println("Starting WebCam Program")
	a.lights.Open()
	a.lights.Reset()
	
	a.WebCamApp  = createWebCamCommand()
}

func (a *app) InitStorageManager() error {
	a.WebCamApp.storageManager.SetWorkDir("/tmp")
	if err := a.WebCamApp.storageManager.Init(); err != nil {
		a.lights.Error()
		return err
	}
	return nil
}

func (a *app) Start() error {
	for {
		a.runner = New(appTimeout)
		a.runner.add(a.WebCamApp)
		
		err := a.runner.Start()
		
		if err != nil && err.Error() != "completed" {
			a.lights.Error()
			return err
		}
	}
	
	return nil
}

func (a *app) Close() {
	a.lights.Reset()
	a.lights.Close()
}

var myapp app = app {
	
}

func main() {
	myapp.lights.SetGPIO(warning.RpioImpl{})
	myapp.Init()
	
	defer myapp.Close()
	
	if err := myapp.InitStorageManager(); err != nil {
		log.Fatal(err)
	}

	myapp.lights.Ok()
	
	if err := myapp.Start(); err != nil {
		log.Fatal(err)
	}
}

func createWebCamCommand() *CameraCommandImpl {
	return &CameraCommandImpl {
		command: "/usr/bin/raspivid",
		args: []string{},
		storageManager: storageManager.New(),
		exec: exec.Command,
	}
}