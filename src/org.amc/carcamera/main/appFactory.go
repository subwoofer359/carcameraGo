package main

import (
	"os"

	"org.amc/carcamera/config"
)

//GetNewApplication factory method for an Application struct
func GetNewApplication(filename string) (app, error) {
	thisApp := app{}

	if tempContext, err := loadConfiguration(filename); err == nil {

		context = tempContext
		thisApp.Init()
		return thisApp, nil
	} else {
		return thisApp, err
	}
}

func loadConfiguration(filename string) (map[string]interface{}, error) {
	var (
		tempContext map[string]interface{}
		err         error
	)

	file, err := os.Open(filename)

	defer file.Close()

	if err != nil {
		return tempContext, err
	}
	if tempContext, err = config.ReadConfigurationFile(file); err != nil {
		return tempContext, err
	}

	return tempContext, nil

}
