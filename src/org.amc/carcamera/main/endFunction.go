package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

type EndCmd interface {
	Stop()
}

type endCmdImpl struct {
	shutdownCMD []string
	syncCMD     []string
}

func GetEndCmdImpl() EndCmd {
	return &endCmdImpl{
		shutdownCMD: []string{"/usr/bin/sudo", "/sbin/shutdown", "-h", "+1"},
		syncCMD:     []string{"/bin/sync"},
	}
}

func (e *endCmdImpl) Stop() {
	syncCmd := exec.Command(e.syncCMD[0])
	err := syncCmd.Run()
	log.Println(err)

	shutdownCmd := exec.Command(e.shutdownCMD[0], e.shutdownCMD[1:]...)
	var stdout, stderr bytes.Buffer
	shutdownCmd.Stdout = &stdout
	shutdownCmd.Stderr = &stderr

	err = shutdownCmd.Run()
	log.Println(err)
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
}
