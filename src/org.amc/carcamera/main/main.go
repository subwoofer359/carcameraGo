package main

import (
	"org.amc/carcamera/warning"
	"log"
)


var myapp app = app {} //myapp Application object

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