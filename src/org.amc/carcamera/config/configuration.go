package config

import (
	//C "org.amc/carcamera/constants"
	"encoding/json"
	"io"
	"log"
	"reflect"
)

type (
	context struct {
		COMMAND string `json:"COMMAND"`
		WORKDIR string `json:"WORKDIR"`
		TIMEOUT string `json:"TIMEOUT"`
		VIDEOLENGTH string `json:"VIDEOLENGTH"`
		PREFIX string `json:"PREFIX"`
		SUFFIX string `json:"SUFFIX"`
		MINFILESIZE string `json:"MINFILESIZE"`
		MAXNOOFFILES string `json:"MAXNOOFFILES`
	}
)

func readConfigurationFile(jsonReader io.Reader) (map[int] string, error) {
	var (
		c context
	)
	
	
	err := json.NewDecoder(jsonReader).Decode(&c)
	
	if err != nil {
		log.Println("Decoding Error")
		return nil, err
	}
	
	config := make(map[int] string)
	
	val := reflect.Indirect(reflect.ValueOf(c))
	for i := 0 ; i < val.NumField(); i = i + 1 {
		log.Println(val.Type().Field(i).Name)	
	}
	
	return config, nil
}