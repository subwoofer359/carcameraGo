package main

import (
	"org.amc/carcamera/warning"
	"org.amc/carcamera/userupdate"
	"log"
	"flag"
)


var (
	myapp app = app {} //myapp Application object
	filename = flag.String("c","" , "Configuration file path")
)

func main() {
	flag.Parse()
	
	if *filename == "" {
		log.Fatal("No configuraton file specified in the command arguments")
	}
	
	if err := myapp.LoadConfiguration(*filename); err != nil {
		log.Fatal(err)
	}
	
	ledService := new(userupdate.LEDService)
	ledService.SetGPIO(warning.RpioImpl{})
	
	myapp.Init()
	
	myapp.message.AddService(ledService)
	
	if err := myapp.message.Init(); err !=nil {
		log.Fatal(err)
	}
	
	defer myapp.Close()
	
	if err := myapp.InitStorageManager(); err != nil {
		log.Fatal(err)
	}

	myapp.message.Started()
	
	if err := myapp.Start(); err != nil {
		log.Fatal(err)
	}
}