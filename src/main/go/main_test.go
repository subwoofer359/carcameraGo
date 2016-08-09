package main

import (
    "testing"
)

func TestRunCommand(t *testing.T) {
	t.Log("Testing web app path");
	command := WebCamApp;
	if(command.Path != "/usr/bin/fswebcam") {
		t.Fatal("incorrect web cam app path");
	}
	t.Log("Testing web app working directory");
	if(command.Dir != "") {
		t.Fatal("incorrect web cam app working directory");
	}
	
}

