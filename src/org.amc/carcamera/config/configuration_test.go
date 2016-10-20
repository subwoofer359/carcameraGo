package config

import (
	C "org.amc/carcamera/constants"
	"testing"
	"strings"
)

var JSON = `{
	"COMMAND": "/bin/ls",
	"WORKDIR": "/tmp",
	"TIMEOUT": "5s",
	"VIDEOLENGTH": "1000",
	"PREFIX": "myPrefix",
	"SUFFIX": "mySuffix",
	"MINFILESIZE": "0",
	"MAXNOOFFILES": "2"
}`

func TestReadConfFile(t *testing.T) {
	
	reader := strings.NewReader(JSON)
	
	context, err := readConfigurationFile(reader)
	
	if err != nil {
		t.Error(err)
	}
	
	if context == nil {
		t.Error("Configuration object wasn't created")
	}
	
	if context[C.COMMAND] != "/bin/ls" {
		t.Error("Configuration not loaded in map")
	}
}