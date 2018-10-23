package main

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
)

//EndCmd commands to be run before application terminates
type EndCmd interface {
	Run()
}

type endCmdImpl struct {
	shutdownCMD []string
	syncCMD     []string
}

//GetEndCmdImpl factory method to create default EndCmd
func GetEndCmdImpl() EndCmd {
	return &endCmdImpl{
		shutdownCMD: []string{"/usr/bin/sudo", "/sbin/shutdown", "-h", "+1"},
		syncCMD:     []string{"/bin/sync"},
	}
}

func (e *endCmdImpl) Run() {
	syncCmd := exec.Command(e.syncCMD[0])
	e.runCmd(syncCmd)

	shutdownCmd := exec.Command(e.shutdownCMD[0], e.shutdownCMD[1:]...)

	e.runCmd(shutdownCmd)
}

func (e *endCmdImpl) runCmd(cmd *exec.Cmd) {
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	e.logError(err)

	errStr := string(stderr.Bytes())

	if len(errStr) > 0 {
		e.logError(errors.New(errStr))
	}

}

func (e *endCmdImpl) logError(err error) {
	if err != nil {
		log.Println(err)
	}
}
