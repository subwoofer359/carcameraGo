package runner

import (
	"log"
	"os"
	"os/exec"
	"testing"
	"time"
)

const timeout = 1 * time.Second

func getNewRunner(d time.Duration) Runner {

	factory := new(SimpleRunnerFactory)

	return factory.NewRunner(d)
}

func TestRunnerStart(t *testing.T) {

	myRunner := getNewRunner(timeout)

	command := &CameraCommandImpl{
		command:        "/bin/ls",
		args:           []string{"/tmp"},
		storageManager: GetMockStorageManagerLS(),
		exec:           exec.Command,
	}

	myRunner.Add(command)
	err := myRunner.Start()

	if err.Error() != "completed" {
		t.Fatal("Error stating runner")
	}
}

func TestRunnerStartTimeOut(t *testing.T) {

	myRunner := getNewRunner(timeout)

	command := &CameraCommandImpl{
		command:        "/bin/dd",
		args:           []string{"if=/dev/urandom"},
		storageManager: GetMockStorageManagerDD(),
		exec:           exec.Command,
	}

	myRunner.Add(command)
	err := myRunner.Start()

	if err.Error() != "received timeout" {
		t.Fatal("Error runner should timeout")
	}
}

func TestRunnerStartInterrupted(t *testing.T) {

	myRunner := RunnerImpl{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(10 * time.Second),
	}

	command := &CameraCommandImpl{
		command:        "/bin/dd",
		args:           []string{"if=/dev/urandom"},
		storageManager: GetMockStorageManagerDD(),
		exec:           exec.Command,
	}

	myRunner.Add(command)

	go interruptCommand(myRunner)

	err := myRunner.Start()

	if err.Error() != "received interrupt" {
		t.Fatal("Error runner should interupt")
	} else {
		log.Println("Process has been interrupted")
	}
}

func interruptCommand(runner RunnerImpl) {
	time.Sleep(5 * time.Millisecond)
	log.Println("Trying to interrupt command")
	runner.interrupt <- os.Interrupt
}
