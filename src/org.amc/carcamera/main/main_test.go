package main

import (
    "testing"
    //"log"
    "time"

)

func TestCameraCommandInitialised(t *testing.T) {
	
	command := WebCamApp
	
	if command.command == "" {
		t.Fatal("Command Not Set")
	}
	
	if command.loop == nil || !command.loop.IsSet() {
		t.Fatal("Loop not set to true which is the default position")
	}
}

func TestRunCommandLoop(t *testing.T) {
	wg.Add(1)
	
	command := WebCamApp
	command.command = "/bin/ls"
	command.args = [] string{"/tmp"}
	
	go command.Run()
	
	time.Sleep(1 * time.Second);
	
	command.loop.UnSet()
	
	wg.Wait()

}

func TestRunCommandLoopThrowsError(t *testing.T) {
	wg.Add(1)
	
	command := WebCamApp
	command.command = "/bin/l"
	command.args = [] string{"/tmp"}
	
	go command.Run()
	
	time.Sleep(1 * time.Second);
	
	command.loop.UnSet()
	
	wg.Wait()

}
