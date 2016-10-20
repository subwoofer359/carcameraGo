package config

import (
	"encoding/json"
	"io"
	"log"
)

func readConfigurationFile(jsonReader io.Reader) (map[string] interface{}, error) {
	var context map[string] interface{}
	
	err := json.NewDecoder(jsonReader).Decode(&context)
	
	if err != nil {
		log.Println("Decoding Error")
		return nil, err
	}
	
	return context, nil
}