package main

import (
	"org.amc/carcamera/warning"
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