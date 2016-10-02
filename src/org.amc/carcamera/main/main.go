package main

import (
	"os/exec"
	"log"
	"sync"
	"github.com/tevino/abool"
	"org.amc/carcamera/storageManager"
)

type CameraCommand struct {
	command string
	args []string
	loop *abool.AtomicBool
	storageManager* storageManager.StorageManager
}

var wg sync.WaitGroup;

func (c *CameraCommand) Run() {
	defer wg.Done()
	
	for c.loop.IsSet() {		
		buffer, err := exec.Command(c.command, c.args...).CombinedOutput();
		if err != nil {
			log.Fatal(err) 
		}
		log.Printf("%s", buffer)
	}
}

var WebCamApp CameraCommand;

func init() {
	WebCamApp = CameraCommand {
		command: "/usr/bin/raspivid",
		args: []string{},
		loop: abool.New(),
		storageManager: storageManager.New(),
	}
	WebCamApp.loop.Set()
	
	WebCamApp.storageManager.WorkDir = "/tmp"
	WebCamApp.storageManager.Init()
}


func main() {
}
