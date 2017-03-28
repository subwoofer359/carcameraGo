package util

import (
	"log"
)

func StringChop(sentence string, length int) []string {
	s := []string{}
	if len(sentence) < length {
		return []string{sentence}
	} else {
		for pointer := 0; pointer < len(sentence); pointer += length {
			end := pointer + length
			if end >= len(sentence){
				end = len(sentence)
			}
			log.Println(sentence[ pointer : end])
			s = append(s,sentence[pointer : end])
		} 
	}
	return s
}