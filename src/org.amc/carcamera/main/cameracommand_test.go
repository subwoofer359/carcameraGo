package main

import (
	"testing"
	"org.amc/carcamera/storageManager"
	
)

func TestRun(t *testing.T) {
	command := CameraCommand{
		command: "/bin/ls",
		args: []string{"/tmp"},
		storageManager: storageManager.New(),		
	}
	
	err := command.Run()
	
	if err.Error() != "completed" {
		t.Fatal("Error running command")
	}
}

func TestRunError(t *testing.T) {
	command := CameraCommand{
		command: "/bin/l",
		args: []string{"/tmp"},
		storageManager: storageManager.New(),		
	}
	
	err := command.Run()
	
	if err.Error() == "completed" {
		t.Fatal("Should have been an Error thrown")
	}
}