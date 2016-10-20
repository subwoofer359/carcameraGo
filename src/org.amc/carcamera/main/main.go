package main

import (
	C "org.amc/carcamera/constants"
	"org.amc/carcamera/storageManager"
	"org.amc/carcamera/warning"
	"os/exec"
	"time"
	"log"
)


var ( 
	myapp app = app {} //myapp Application object
	context = map[string] string {
		C.COMMAND: "/usr/bin/raspivid",
		C.WORKDIR: "/mnt/external",
		C.TIMEOUT: "7m",
		C.VIDEOLENGTH: "300000",
		C.PREFIX: "video",
		C.SUFFIX: ".h264",
		C.MINFILESIZE: "0",
		C.MAXNOOFFILES: "20",
		
	}
)

type app struct {
	runner *Runner
	lights warning.UserDisplay
	WebCamApp *CameraCommandImpl
	appTimeOut time.Duration
}

func (a *app) Init() {
	log.Println("Starting WebCam Program")
	a.lights.Open()
	a.lights.Reset()
	
	a.WebCamApp  = createWebCamCommand()
	
	timeout,_ := time.ParseDuration(context[C.TIMEOUT])
	
	a.appTimeOut = timeout
}

func createWebCamCommand() *CameraCommandImpl {
	return &CameraCommandImpl {
		command: context[C.COMMAND],
		args: []string{"-t", context[C.VIDEOLENGTH], "-rot", "270", "-o"},
		storageManager: storageManager.New(context),
		exec: exec.Command,
	}
}

func (a *app) InitStorageManager() error {
	a.WebCamApp.storageManager.SetWorkDir(context[C.WORKDIR])
	if err := a.WebCamApp.storageManager.Init(); err != nil {
		a.lights.Error()
		return err
	}
	return nil
}

func (a *app) Start() error {
	for {
		a.runner = New(a.appTimeOut)
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