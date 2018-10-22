package main

import (
	"os"
	"testing"

	C "org.amc/carcamera/constants"
)

var testFilename = os.ExpandEnv("$GOPATH") + "/configuration.json"

func TestLoadConfiguration(t *testing.T) {

	testContext, err := loadConfiguration(testFilename)

	if err != nil {
		t.Error(err)
	}

	if testContext[C.COMMAND] == nil || testContext[C.COMMAND] == "" {
		t.Errorf("%s shouldn't be empty\n", C.COMMAND)
	}
}

func TestGetNewApplicationFileError(t *testing.T) {
	if _, err := GetNewApplication("/tmp"); err == nil {
		t.Error("Error should have been thrown")
	}
}

func TestGetNewApplicationInitialised(t *testing.T) {

	var (
		testApp app
		err     error
	)
	if testApp, err = GetNewApplication(testFilename); err == nil {
		if testApp.runnerFactory == nil {
			t.Error("Application not initialised")
		}

		if context[C.COMMAND] == nil {
			t.Error("Context not set")
		}
	} else {
		t.Error(err)
	}

}
