package config

import (
	C "org.amc/carcamera/constants"
	"strings"
	"testing"
)

var (
	JSON = `{
		"COMMAND": "/bin/ls",
		"WORKDIR": "/tmp",
		"TIMEOUT": "5s",
		"VIDEOLENGTH": "1000",
		"PREFIX": "myPrefix",
		"SUFFIX": "mySuffix",
		"MINFILESIZE": "0",
		"MAXNOOFFILES": "2",
		"OPTIONS": "-rot 90"
	}`

	JSON_NO_OPTIONS = `{
		"COMMAND": "/bin/ls",
		"WORKDIR": "/tmp",
		"TIMEOUT": "5s",
		"VIDEOLENGTH": "1000",
		"PREFIX": "myPrefix",
		"SUFFIX": "mySuffix",
		"MINFILESIZE": "0",
		"MAXNOOFFILES": "2"
	}`
)

func TestReadConfFile(t *testing.T) {
	reader := strings.NewReader(JSON)

	context, err := ReadConfigurationFile(reader)

	if err != nil {
		t.Error(err)
	}

	if context == nil {
		t.Error("Configuration object wasn't created")
	}

	if context[C.COMMAND] != "/bin/ls" {
		t.Error("Configuration not loaded in map")
	}

	if context[C.MAXNOOFFILES] != "2" {
		t.Errorf("Configuration not loaded in map for %s\n", C.MAXNOOFFILES)
	}
	options := context[C.OPTIONS].([]string)

	optionStr := strings.Join(options, " ")

	if optionStr != "-t 1000 -rot 90" {
		t.Errorf("Options string not parsed correctly:%s", optionStr)
	}

}

func TestReadConfFileNoOptionTag(t *testing.T) {
	reader := strings.NewReader(JSON_NO_OPTIONS)

	context, err := ReadConfigurationFile(reader)

	if err != nil {
		t.Error(err)
	}

	if context == nil {
		t.Error("Configuration object wasn't created")
	}

	if context[C.OPTIONS] != nil {
		t.Error("Options shouldn't be set")
	}
}
