package config

import (
	C "org.amc/carcamera/constants"
	"encoding/json"
	"io"
	"log"
	"strings"
)

func ReadConfigurationFile(jsonReader io.Reader) (map[string] interface{}, error) {
	var context map[string] interface{}
	
	err := json.NewDecoder(jsonReader).Decode(&context)
	
	if err != nil {
		log.Println("Decoding Error")
		return nil, err
	}
	//concatenate options string and durations string which are separate
	log.Println(context[C.VIDEOLENGTH].(string))
	options := append([]string {"-t", context[C.VIDEOLENGTH].(string)}, 
		strings.Split(context[C.OPTIONS].(string), " ")...)
	context[C.OPTIONS] = options
	return context, nil
}