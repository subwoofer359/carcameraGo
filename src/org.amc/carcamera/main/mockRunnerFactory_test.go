package main

import (
	"errors"
	"log"
	"time"

	"org.amc/carcamera/runner"
)

//runnerCalls to keep track of calls to runner.Start()
var runnerCalls int

type mockRunner struct{}

func (m *mockRunner) Start() error {
	log.Printf("M Calls:%d", runnerCalls)
	if runnerCalls > 5 {
		return errTestStopped
	}
	runnerCalls++
	return errors.New(runner.COMPLETED)
}

func (m mockRunner) Stop() {}

func (m mockRunner) Handle() error {
	return nil
}

func (m *mockRunner) Add(command runner.CameraCommand) {}

type mockRunnerFactory struct{}

func (m mockRunnerFactory) NewRunner(d time.Duration) runner.Runner {
	return &mockRunner{}
}

// A Mock Runner that executes slowly
type slowMockRunnerFactory struct{}

func (m slowMockRunnerFactory) NewRunner(d time.Duration) runner.Runner {
	return &slowMockRunner{}
}

type slowMockRunner struct {
	mockRunner
}

func (m *slowMockRunner) Start() error {
	time.Sleep(1 * time.Second)
	return m.mockRunner.Start()
}
