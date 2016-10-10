package storageManager

import (
	"strconv"
	"io/ioutil"
	"log"
	"regexp"
	"os"
)

//PREFIX Filename prefix
const PREFIX = "video"

//SUFFIX Filename suffix
const SUFFIX = ".mpg" 

//const MIN_FILE_SIZE = 1024 * 1000

//MinFileSize The minimum file size to be accepted
const MinFileSize = 0 

//MaxNoOfFiles The maximum number of files
const MaxNoOfFiles = 20 

//StorageManager object
type StorageManager struct {
	prefix  string
	suffix  string
	index 	int
	WorkDir string
	fileList []string
}

//New create new StorageManager
func New() *StorageManager {
	s := new(StorageManager)
	return s
}

//MaxNoOfFiles return MaxNoOfFiles
func (s StorageManager) MaxNoOfFiles() int {
	return MaxNoOfFiles
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
	for len(s.fileList) > MaxNoOfFiles {
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
	if file.Size() > MinFileSize {
		s.fileList = append(s.fileList, fileName)
		removeOldFiles(s)
		return nil
	}
	
	return os.Remove(fileName)
}

