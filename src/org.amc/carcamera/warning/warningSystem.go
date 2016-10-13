package warning

import (
	gpio "github.com/stianeikeland/go-rpio"
	"log"	
)

func Yellow() {
	err := gpio.Open()
	
	if err != nil {
		log.Fatal(err)
	}
	
	pin := gpio.Pin(10)
	
	pin.Output()
	
	pin.High()
	
	defer gpio.Close()
}