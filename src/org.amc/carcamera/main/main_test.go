package main

import (
    "testing"
    //"log"
    //"time"

)

func TestCreateCameraCommand(t *testing.T) {
	command := createWebCamCommand()
	
	if command == nil {
		t.Error("Command not created")
	}
}
