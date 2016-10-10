package main

import (
	"os/exec"
	"sync"
	"errors"
	"strings"
	"log"
	"org.amc/carcamera/storageManager"
)

var wg sync.WaitGroup

type CameraCommand struct {
	command string
	args []string
	storageManager* storageManager.StorageManager
}

func (c *CameraCommand) Run() error {
	defer wg.Done()
				
	buffer, err := exec.Command(c.command, c.args...).CombinedOutput();
	if err != nil {
		return errors.New(strings.Join([]string{"ERROR:", err.Error()}, " "))

	} 
	log.Printf("%s", buffer)
	
	return errors.New("completed")
}
