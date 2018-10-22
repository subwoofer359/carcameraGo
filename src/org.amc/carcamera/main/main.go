package main

import (
	"flag"
	"log"
	"time"

	C "org.amc/carcamera/constants"
	"org.amc/carcamera/userupdate"
	"org.amc/carcamera/warning"
)

var (
	myapp    = app{} //myapp Application object
	filename = flag.String("c", "", "Configuration file path")
	appGpio  = warning.RpioImpl{}
)

func main() {
	flag.Parse()

	if *filename == "" {
		log.Fatal("No configuraton file specified in the command arguments")
	}

	if err := myapp.LoadConfiguration(*filename); err != nil {
		log.Fatal(err)
	}

	myapp.Init()

	myapp.endCmd = GetEndCmdImpl()

	addPowerControl()

	setUpServices()

	if err := myapp.message.Init(); err != nil {
		log.Fatal(err)
	}

	defer myapp.Close()
	myapp.message.Started()

	if err := myapp.InitStorageManager(); err != nil {
		myapp.message.Error(err.Error())
		mainExit()
	} else if err := myapp.Start(); err != nil {
		myapp.message.Error(err.Error())
		mainExit()
	}
}

func addPowerControl() {
	pc := &PowerControlImpl{
		gpio: appGpio,
	}

	if err := pc.Init(); err != nil {
		log.Println(err)
		myapp.message.Error(err.Error())
		mainExit()
	}

	myapp.powerControl = pc.poweroff
	pc.Start()
}

func setUpServices() {
	if context[C.BLUETOOTH] != C.OFF {
		btService := new(userupdate.BTService)
		btService.SetContext(context)
		myapp.message.AddService(btService)
	}

	ledService := new(userupdate.LEDService)
	ledService.SetGPIO(appGpio)
	myapp.message.AddService(ledService)

}

func mainExit() {
	myapp.message.Stopped()
	time.Sleep(30 * time.Second)
	log.Println("Program stopped due to error conditions")
}
