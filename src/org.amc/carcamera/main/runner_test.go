package main

import (
	"testing"
	"time"
	"os"
	"log"
	"org.amc/carcamera/storageManager"
)

const timeout = 1 * time.Second 

func TestRunnerStart(t *testing.T) {
	
	myRunner := New(timeout);
	
	command := &CameraCommand{
		command: "/bin/ls",
		args: []string{"/tmp"},
		storageManager: storageManager.New(),
	}
	
	myRunner.add(command)
	err := myRunner.Start()
	
	if err.Error() != "completed" {
		t.Fatal("Error stating runner")
	}
}


func TestRunnerStartTimeOut(t *testing.T) { 
	
	myRunner := New(timeout);
	
	command := &CameraCommand{
		command: "/bin/dd",
		args: []string{"if=/dev/urandom", "of=/tmp/e"},
		storageManager: storageManager.New(),
	}
	
	myRunner.add(command)
	err := myRunner.Start()
	
	if err.Error() != "received timeout" {
		t.Fatal("Error runner should timeout")
	}
}

func TestRunnerStartInterrupted(t *testing.T) { 
	
	myRunner := New(10 * time.Second);
	
	command := &CameraCommand{
		command: "/bin/dd",
		args: []string{"if=/dev/urandom", "of=/tmp/e"},
		storageManager: storageManager.New(),
	}
	
	myRunner.add(command)
	
	go interruptCommand(myRunner)
	
	err := myRunner.Start()
	
	if err.Error() != "received interrupt" {
		t.Fatal("Error runner should interupt")
	} else {
		log.Println("Process has been interrupted")
	}
}

func interruptCommand(runner *Runner) {
	time.Sleep(2 * time.Second)
	log.Println("Trying to interrupt command")	
	runner.interrupt <- os.Interrupt
}
