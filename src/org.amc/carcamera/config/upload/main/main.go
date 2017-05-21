package main

import (
	"flag"
	"log"
	"os"
	"os/exec"

	"org.amc/carcamera/config"
	"org.amc/carcamera/config/upload"
)

var (
	filenameOption = flag.String("c", "", "Configuration file path")
	dashcamName    = "dashcam"
	configFile     = "configuration.json"
)

func main() {
	context := loadConfigurationFile()

	if currentPath, err := os.Getwd(); err == nil {
		log.Println("Attempting to update executable")
		updateExecutable(context, currentPath)

		log.Println("Attempting to update configuration file")
		updateConfigFile(context, currentPath)
	}
}

func loadConfigurationFile() map[string]interface{} {
	var (
		file    *os.File
		err     error
		context map[string]interface{}
	)

	flag.Parse()

	if *filenameOption == "" {
		log.Fatal("No configuraton file specified in the command arguments")
	}

	if file, err = os.Open(*filenameOption); err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if context, err = config.ReadConfigurationFile(file); err != nil {
		log.Fatal(err)
	}

	return context
}

func updateExecutable(context map[string]interface{}, currentPath string) {
	updater := upload.NewUpdater(context, currentPath, dashcamName)
	if err := updater.UpdateExecutable(); err != nil {
		log.Fatal(err)
	} else {
		setcap(dashcamName)
	}
}

func updateConfigFile(context map[string]interface{}, currentPath string) {
	updater := upload.NewUpdater(context, currentPath, configFile)
	if err := updater.UpdateExecutable(); err != nil {
		log.Fatal(err)
	}
}

//setcap set root permissions on the executable which is needed for
// bluetooth control
func setcap(filename string) {
	log.Println("SetCap on " + filename)
	cmd := exec.Command("/bin/sh", "-c", "sudo setcap 'cap_net_raw,cap_net_admin=eip' "+filename)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
