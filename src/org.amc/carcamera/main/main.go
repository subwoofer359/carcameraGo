package main

import (
	"org.amc/carcamera/warning"
	"org.amc/carcamera/userupdate"
	"log"
	"flag"
	"time"
)


var (
	myapp app = app {} //myapp Application object
	filename = flag.String("c", "", "Configuration file path")
)

func main() {
	flag.Parse()
	
	if *filename == "" {
		log.Fatal("No configuraton file specified in the command arguments")
	}
	
	if err := myapp.LoadConfiguration(*filename); err != nil {
		log.Fatal(err)
	}
	
	btService := new(userupdate.BTService)
	
	ledService := new(userupdate.LEDService)
	ledService.SetGPIO(warning.RpioImpl{})
	
	myapp.Init()
	
	btService.SetContext(context)
	myapp.message.AddService(ledService)
	myapp.message.AddService(btService)
	
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

func mainExit() {
        myapp.message.Stopped()
        time.Sleep(30 * time.Second)
        log.Fatal("Program stopped due to error conditions")
}
