package main

import (
	"log"
	"os/exec"
	"strings"
	"testing"
)

func TestCameraCommandRun(t *testing.T) {
	command := CameraCommandImpl{
		command:        "/bin/ls",
		storageManager: GetMockStorageManagerLS(),
		exec:           exec.Command,
	}

	err := command.Run()

	if err.Error() != "completed" {
		t.Fatal("Error running command")
	}
}

func TestCameraCommandRunError(t *testing.T) {
	command := CameraCommandImpl{
		command:        "/bin/l",
		storageManager: GetMockStorageManagerLS(),
		exec:           exec.Command,
	}

	err := command.Run()

	if err.Error() == "completed" {
		t.Fatal("Should have been an Error thrown")
	}
}

func TestStdoutPipeError(t *testing.T) {

	newCmd := exec.Cmd{}

	s, _ := newCmd.StdoutPipe()
	s.Close()

	err := runPipeTest(&newCmd)

	if err == nil {
		t.Fatal(err)
	}
}

func runPipeTest(newCmd *exec.Cmd) error {
	command := CameraCommandImpl{
		command:        "/bin/l",
		storageManager: GetMockStorageManagerLS(),
		exec:           func(name string, args ...string) *exec.Cmd { return newCmd },
	}

	return command.Run()
}

func TestStderrPipeError(t *testing.T) {

	newCmd := exec.Cmd{}

	s, _ := newCmd.StderrPipe()
	s.Close()

	err := runPipeTest(&newCmd)

	if err == nil {
		t.Fatal(err)
	}
}

func TestStderrPipe(t *testing.T) {
	errOutput := "There has been a terrible failure detected"

	reader := strings.NewReader(errOutput)

	err := readErrPipe(reader)

	if err == nil {
		t.Fatal("An error should be returned from pipe read")
	}

	if err.Error() != errOutput {
		t.Fatalf("Error message should be (%s) not (%s)", errOutput, err.Error())
	} else {
		log.Printf("Error message (%s) was returned which is correct", err.Error())
	}
}
