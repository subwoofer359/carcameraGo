package storageManager

import (
	"strconv"
	"io/ioutil"
	"log"
	"regexp"
)
const PREFIX = "video";
const SUFFIX = ".mpg";

type storageManager struct {
	prefix  string
	suffix  string
	index 	int
	workDir string
}

func (s *storageManager) getFileName() string {
	incr := strconv.Itoa(s.index);
	s.index = s.index + 1;
	return PREFIX + incr + SUFFIX
}

func (s *storageManager) setLastIndex() {
	pattern := PREFIX + "(\\d+)\\" + SUFFIX
	matcher := regexp.MustCompile(pattern)
	files, err := ioutil.ReadDir(s.workDir)
	index := 0;
	if err != nil {
		log.Fatal(err)
	} else {
		for _, file := range files {
			matches := matcher.FindStringSubmatch(file.Name())
			if(len(matches) > 0) {
				tmpIndex, _ := strconv.Atoi(matches[1])
				if(tmpIndex > index) {
					index = tmpIndex;
				}		
			}
		}
		
	}
	
	s.index = index;
}

