package main

import (
	"log"
	"os"
	"os/exec"
	"testing"
	"time"
)

const timeout = 1 * time.Second

func TestRunnerStart(t *testing.T) {

	myRunner := NewRunner(timeout)

	command := &CameraCommandImpl{
		command:        "/bin/ls",
		args:           []string{"/tmp"},
		storageManager: GetMockStorageManagerLS(),
		exec:           exec.Command,
	}

	myRunner.add(command)
	err := myRunner.Start()

	if err.Error() != "completed" {
		t.Fatal("Error stating runner")
	}
}

func TestRunnerStartTimeOut(t *testing.T) {

	myRunner := NewRunner(timeout)

	command := &CameraCommandImpl{
		command:        "/bin/dd",
		args:           []string{"if=/dev/urandom"},
		storageManager: GetMockStorageManagerDD(),
		exec:           exec.Command,
	}

	myRunner.add(command)
	err := myRunner.Start()

	if err.Error() != "received timeout" {
		t.Fatal("Error runner should timeout")
	}
}

func TestRunnerStartInterrupted(t *testing.T) {

	myRunner := NewRunner(10 * time.Second)

	command := &CameraCommandImpl{
		command:        "/bin/dd",
		args:           []string{"if=/dev/urandom"},
		storageManager: GetMockStorageManagerDD(),
		exec:           exec.Command,
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
