package main

import (
	"org.amc/carcamera/config"
	C "org.amc/carcamera/constants"
	"org.amc/carcamera/warning"
	"org.amc/carcamera/storageManager"
	"log"
	"os"
	"os/exec"
	"time"
)

var	context map[string] string

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

func (a app) LoadConfiguration(filename string) error {
	file, err := os.Open(filename)
	defer file.Close()
	
	if err != nil {
		return err
	}
	if tempContext, err := config.ReadConfigurationFile(file); err != nil {
		return err
	} else {
		context = tempContext
	}
	return nil
	
}