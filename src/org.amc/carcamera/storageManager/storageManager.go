package storageManager

import (
	"strconv"
	"io/ioutil"
	"log"
	"regexp"
	"os"
)
const PREFIX = "video";
const SUFFIX = ".mpg";

/* 
 *The minimum file size to be accepted
 */
//const MIN_FILE_SIZE = 1024 * 1000
const MIN_FILE_SIZE = 0

/*
 * The maximum number of files
 */
const MAX_NO_OF_FILES = 20


type StorageManager struct {
	prefix  string
	suffix  string
	index 	int
	WorkDir string
	fileList []string
}

func New() *StorageManager {
	s := new(StorageManager)
	return s
}

func (s StorageManager) MaxNoOfFiles() int {
	return MAX_NO_OF_FILES
}

func (s *StorageManager) Init() {
	log.Println("StorageManager Init called")
	
	s.index, s.fileList = findAndSaveExistingFileNames(s.WorkDir);
}

func findAndSaveExistingFileNames(workDir string) (int, []string) {
	pattern := PREFIX + "(\\d+)\\" + SUFFIX
	matcher := regexp.MustCompile(pattern)
	files, err := ioutil.ReadDir(workDir)
	index := 0;
	maxIndex := 0;
	fileList := []string{};
	
	if err != nil {
		log.Fatal(err)
	} else {
		for _, file := range files {
			matches := matcher.FindStringSubmatch(file.Name())
			if(len(matches) > 0) {
				fileList = append(fileList, workDir + "/" + file.Name())
				tmpIndex, _ := strconv.Atoi(matches[1])
				if(tmpIndex > index) {
					index = tmpIndex;
				}		
			}
		}
		maxIndex = index;	
	}
	return maxIndex, fileList;
}

func (s *StorageManager) GetNextFileName() string {
	incr := strconv.Itoa(s.index);
	s.index = s.index + 1;
	
	newFileName := s.WorkDir + "/" + PREFIX + incr + SUFFIX;
	
	return newFileName
}

func removeOldFiles(s *StorageManager) {
	for len(s.fileList) > MAX_NO_OF_FILES {
		s.RemoveLRU()
	}
}

func (s *StorageManager) RemoveLRU() {
	if len(s.fileList) > 0 {
		oldFileStr := s.fileList[0]
		err := os.Remove(oldFileStr)
		if err != nil {
			log.Println(err)
		}
		s.fileList = s.fileList[1:]
		
	}
}

func (s *StorageManager) addCompleteFile(fileName string) error {
	
	file, err := os.Stat(fileName);
	if err != nil {
		return err
	}
	if file.Size() > MIN_FILE_SIZE {
		s.fileList = append(s.fileList, fileName)
		removeOldFiles(s)
		return nil
	} else {
		return os.Remove(fileName)
	}
}

