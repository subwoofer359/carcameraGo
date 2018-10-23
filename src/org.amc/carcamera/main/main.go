package main

import (
	"flag"
	"log"
	"time"

	C "org.amc/carcamera/constants"
	"org.amc/carcamera/powercontrol"
	"org.amc/carcamera/userupdate"
	"org.amc/carcamera/warning"
)

var (
	myapp    app //myapp Application object
	filename = flag.String("c", "", "Configuration file path")
	appGpio  = warning.RpioImpl{}
)

func main() {
	var err error
	flag.Parse()

	if *filename == "" {
		log.Fatal("No configuraton file specified in the command arguments")
	}

	if myapp, err = GetNewApplication(*filename); err != nil {
		log.Fatal(err)
	}

	defer myapp.Close()

	setUpApplication()

	startApplication()
}

func setUpApplication() {
	myapp.endCmd = GetEndCmdImpl()

	addPowerControl()

	setUpServices()

	if err := myapp.message.Init(); err != nil {
		log.Fatal(err)
	}

	myapp.message.Started()
}

func addPowerControl() {
	pc := &powercontrol.PowerControlImpl{}

	pc.SetGPIO(appGpio)

	if err := pc.Init(); err != nil {
		log.Println(err)
		myapp.message.Error(err.Error())
		mainExit()
	}

	myapp.powerControl = pc.PowerOff()
	pc.Start()
}

func setUpServices() {
	//Add Console Logging
	consoleService := new(userupdate.ConsoleService)
	myapp.message.AddService(consoleService)

	//Add Bluetooth logging if set in configuration
	if context[C.BLUETOOTH] != C.OFF {
		btService := new(userupdate.BTService)
		btService.SetContext(context)
		myapp.message.AddService(btService)
	}

	//Add Led lights messaging
	ledService := new(userupdate.LEDService)
	ledService.SetGPIO(appGpio)
	myapp.message.AddService(ledService)

}

func startApplication() {
	if err := myapp.InitStorageManager(); err != nil {
		myapp.message.Error(err.Error())
		mainExit()
	} else if err := myapp.Start(); err != nil {
		myapp.message.Error(err.Error())
		mainExit()
	}
}

func mainExit() {
	myapp.message.Stopped()
	myapp.endCmd.Run()
	time.Sleep(30 * time.Second)
	log.Println("Program stopped due to error conditions")
}
