package main
//
//import (
//	"testing"
//	"org.amc/carcamera/storageManager"
//	"os/exec"
//	"time"
//	"log"
//)
//
//func TestIntegration(t *testing.T) {
//	WebCamApp := &CameraCommandImpl {
//		command: "/usr/bin/avconv",
//		args: []string{"-f", "video4linux2", "-i", "/dev/video0", "-t", "5", "-y"},
//		storageManager: storageManager.New(),
//		exec: exec.Command,
//	}
//	
//	WebCamApp.storageManager.SetWorkDir("/tmp")
//	WebCamApp.storageManager.Init()
//	
//	
//	
//	for t := 0; t < 2; t++ {
//		runner := New(7 * time.Second)
//		runner.add(WebCamApp)
//		err := runner.Start()
//		if err != nil && err.Error() != "completed" {
//			log.Fatal(err)	
//		}
//	}
//}
//
