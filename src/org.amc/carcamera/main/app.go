package main

import (
	"errors"
	"log"
	"os/exec"
	"time"

	C "org.amc/carcamera/constants"
	"org.amc/carcamera/powercontrol"
	"org.amc/carcamera/runner"
	"org.amc/carcamera/storageManager"
	"org.amc/carcamera/userupdate"
	"org.amc/carcamera/warning"
)

var (
	context        map[string]interface{}
	defaultFactory runner.RunnerFactory = new(runner.SimpleRunnerFactory)
)

type app struct {
	runnerFactory runner.RunnerFactory
	lights        warning.UserDisplay
	message       *userupdate.Message
	WebCamApp     runner.CameraCommand
	appTimeOut    time.Duration
	powerControl  chan bool
	endCmd        EndCmd
}

func (a *app) Init() {
	log.Println("Starting WebCam Program")

	a.message = new(userupdate.Message)

	a.WebCamApp = createWebCamCommand()

	a.appTimeOut, _ = time.ParseDuration(context[C.TIMEOUT].(string))

	if a.runnerFactory == nil {
		a.runnerFactory = defaultFactory
	}
}

func createWebCamCommand() runner.CameraCommand {
	commandFactory := new(runner.SimpleCameraCommandFactory)
	return commandFactory.NewCameraCommand(
		context[C.COMMAND].(string),
		context[C.OPTIONS].([]string),
		storageManager.NewStorageManager(context),
		exec.Command)
}

func (a *app) InitStorageManager() error {
	if err := a.WebCamApp.StorageManager().Init(); err != nil {
		a.message.Error(err.Error())
		return err
	}
	a.WebCamApp.StorageManager().RemoveOldFiles()
	return nil
}

func (a *app) Start() error {
	if a.powerControl == nil {
		return errors.New("PowerControl is nil")
	}

	for {
		var arunner = a.runnerFactory.NewRunner(a.appTimeOut)
		arunner.Add(a.WebCamApp)

		result := make(chan error)

		go func() {
			result <- arunner.Start()
		}()

		select {
		case err := <-result:
			//If  Completed received then keep looping
			if err != nil && err.Error() != runner.COMPLETED {
				a.message.Error(err.Error())

				return err
			}
		case <-a.powerControl:
			arunner.Stop()
			a.endCmd.Run()
			return powercontrol.ErrPowerFault
		}

		close(result)
	}
}

func (a *app) Close() {
	a.message.Stopped()
	a.message.Close()
}
