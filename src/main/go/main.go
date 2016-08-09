package main

import (
	"os/exec"
	"fmt"
)

var WebCamApp = exec.Command("/usr/bin/fswebcam");

func main() {
	buffer, err := WebCamApp.CombinedOutput();
	if err != nil {
		fmt.Println(err.Error());
	}
	
	fmt.Printf("process is %d\n", WebCamApp.Process.Pid);
	
	fmt.Printf("***%s****", buffer);
}
