package main

import (
	"org.amc/carcamera/config"
	C "org.amc/carcamera/constants"
	"org.amc/carcamera/warning"
	"org.amc/carcamera/storageManager"
	"org.amc/carcamera/userupdate"
	"log"
	"os"
	"os/exec"
	"time"
)

var	context map[string] interface{}

type app struct {
	runner *Runner
	lights warning.UserDisplay
	message *userupdate.Message
	WebCamApp *CameraCommandImpl
	appTimeOut time.Duration
}

func (a *app) Init() {
	log.Println("Starting WebCam Program")
	
	a.message = new(userupdate.Message)
	
	a.WebCamApp  = createWebCamCommand()
	
	timeout,_ := time.ParseDuration(context[C.TIMEOUT].(string))
	
	a.appTimeOut = timeout
	
}

func createWebCamCommand() *CameraCommandImpl {
	return &CameraCommandImpl {
		command: context[C.COMMAND].(string),
		args: context[C.OPTIONS].([]string),
		storageManager: storageManager.New(context),
		exec: exec.Command,
	}
}

func (a *app) InitStorageManager() error {
	a.WebCamApp.storageManager.SetWorkDir(context[C.WORKDIR].(string))
	if err := a.WebCamApp.storageManager.Init(); err != nil {
		a.message.Error(err.Error())
		return err
	}
	return nil
}

func (a *app) Start() error {
	for {
		a.runner = NewRunner(a.appTimeOut)
		a.runner.add(a.WebCamApp)
		
		err := a.runner.Start()
		
		if err != nil && err.Error() != "completed" {
			a.message.Error(err.Error())
			return err
		}
	}
	
	return nil
}

func (a *app) Close() {
	a.message.Stopped()
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