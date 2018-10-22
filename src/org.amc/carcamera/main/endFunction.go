package main

import (
	"bytes"
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
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	e.logError(err)

	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	log.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
}

func (e *endCmdImpl) logError(err error) {
	if err != nil {
		log.Println(err)
	}
}
