package main

import (
	"os/exec"
	"errors"
	"strings"
	"log"
	"org.amc/carcamera/storageManager"
)

type CameraCommand struct {
	command string
	args []string
	storageManager* storageManager.StorageManager
}

func (c *CameraCommand) Run() error {
				
	buffer, err := exec.Command(c.command, c.args...).CombinedOutput();
	if err != nil {
		return errors.New(strings.Join([]string{"ERROR:", err.Error()}, " "))

	} 
	log.Printf("%s", buffer)
	
	return errors.New("completed")
}
