package main

import (
    "testing"
    //"log"
    //"time"

)

func TestCameraCommandInitialised(t *testing.T) {
	
	command := WebCamApp
	
	if command.command == "" {
		t.Fatal("Command Not Set")
	}
}
