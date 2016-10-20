package config

import (
	"encoding/json"
	"io"
	"log"
)

func ReadConfigurationFile(jsonReader io.Reader) (map[string] string, error) {
	var context map[string] string
	
	err := json.NewDecoder(jsonReader).Decode(&context)
	
	if err != nil {
		log.Println("Decoding Error")
		return nil, err
	}
	
	return context, nil
}