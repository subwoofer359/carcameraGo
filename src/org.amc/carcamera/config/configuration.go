package config

import (
	"encoding/json"
	"io"
	"log"
	C "org.amc/carcamera/constants"
	"strings"
)

func ReadConfigurationFile(jsonReader io.Reader) (map[string]interface{}, error) {
	var context map[string]interface{}

	err := json.NewDecoder(jsonReader).Decode(&context)

	if err != nil {
		log.Println("Decoding Error")
		return nil, err
	}
	//concatenate options string and durations string which are separate
	log.Println(context[C.VIDEOLENGTH].(string))

	if context[C.OPTIONS] == nil {
		log.Println("Config: No Options set in configuration file")
	} else {
		options := append([]string{"-t", context[C.VIDEOLENGTH].(string)},
			strings.Split(context[C.OPTIONS].(string), " ")...)
		context[C.OPTIONS] = options
	}
	return context, nil
}
