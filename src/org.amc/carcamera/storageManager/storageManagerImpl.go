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
type StorageManagerImpl struct {
	prefix  string
	suffix  string
	index 	int
	workDir string
	fileList []string
}

//New create new StorageManager
func New() StorageManager {
	s := new(StorageManagerImpl)
	return s
}

func (s StorageManagerImpl) Prefix() string {
	return s.prefix;
}

func (s StorageManagerImpl) Suffix() string {
	return s.prefix
}

func (s StorageManagerImpl) Index() int {
	return s.index
}

func (s StorageManagerImpl) WorkDir() string {
	return s.workDir
}

func (s StorageManagerImpl) FileList() []string {
	return s.fileList
}

//MaxNoOfFiles return MaxNoOfFiles
func (s StorageManagerImpl) MaxNoOfFiles() int {
	return MaxNoOfFiles
}

func (s *StorageManagerImpl) Init() {
	log.Println("StorageManager Init called")
	
	s.index, s.fileList = findAndSaveExistingFileNames(s.WorkDir());
	
	s.index = s.index + 1;
	
	log.Printf("StorageManager: %d previous files found\n", s.index)
	log.Println(s.fileList)
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

func (s *StorageManagerImpl) GetNextFileName() string {
	incr := strconv.Itoa(s.index);
	s.index = s.index + 1;
	
	newFileName := s.WorkDir() + "/" + PREFIX + incr + SUFFIX;
	
	return newFileName
}

func removeOldFiles(s *StorageManagerImpl) {
	for len(s.fileList) > MaxNoOfFiles {
		s.RemoveLRU()
	}
}

func (s *StorageManagerImpl) RemoveLRU() {
	if len(s.fileList) > 0 {
		oldFileStr := s.fileList[0]
		err := os.Remove(oldFileStr)
		if err != nil {
			log.Println(err)
		}
		s.fileList = s.fileList[1:]
		
	}
}

func (s *StorageManagerImpl) AddCompleteFile(fileName string) error {
	
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