package storageManager

import (
	"strconv"
	"io/ioutil"
	"log"
	"regexp"
	"os"
	"fmt"
)

//StorageManager object
type StorageManagerImpl struct {
	index 	int
	fileList []string
	context map[string] string
}

//New create new StorageManager
func New(context map[string] string) StorageManager {
	s := new(StorageManagerImpl)
	s.context = context
	return s
}

func (s StorageManagerImpl) Prefix() string {
	return s.context["PREFIX"];
}

func (s StorageManagerImpl) Suffix() string {
	return s.context["SUFFIX"]
}

func (s StorageManagerImpl) Index() int {
	return s.index
}

func (s StorageManagerImpl) WorkDir() string {
	return s.context["WORKDIR"]
}

func (s *StorageManagerImpl) SetWorkDir(workDir string) {
	s.context["WORKDIR"] = workDir
}

func (s StorageManagerImpl) FileList() []string {
	return s.fileList
}

//MaxNoOfFiles return MaxNoOfFiles
func (s StorageManagerImpl) MaxNoOfFiles() int {
	maxfiles,_ := strconv.Atoi(s.context["MAXNOOFFILES"])
	return maxfiles
}

func (s StorageManagerImpl) MinFileSize() int64 {
	minfileSize,_ := strconv.Atoi(s.context["MINFILESIZE"])
	return int64(minfileSize)
}

func (s *StorageManagerImpl) GetContext() map[string] string {
	return s.context
}

func (s *StorageManagerImpl) Init() error {
	log.Println("StorageManager Init called")
	
	if index, fileList , err := findAndSaveExistingFileNames(s); err != nil {
		return fmt.Errorf("Error reading Work Directory %s\n", s.WorkDir())
	} else {
		s.index = index
		s.fileList = fileList
	}
	
	log.Printf("StorageManager: %d previous files found\n", s.index)
	
	s.index = s.index + 1;
	
	log.Printf("StorageManager:%s\n", s.fileList)
	
	return nil
}

func findAndSaveExistingFileNames(s *StorageManagerImpl) (int, []string, error) {
	pattern := s.Prefix() + "(\\d+)\\" + s.Suffix()
	matcher := regexp.MustCompile(pattern)
	maxIndex := 0;
	fileList := []string{};
	
	if files, err := ioutil.ReadDir(s.WorkDir()); err != nil {
		return 0, nil, err
	} else {
		for _, file := range files {
			matches := matcher.FindStringSubmatch(file.Name())
			if(len(matches) > 0) {
				fileList = append(fileList, s.WorkDir() + "/" + file.Name())
				tmpIndex, _ := strconv.Atoi(matches[1])
				if(tmpIndex > maxIndex) {
					maxIndex = tmpIndex;
				}		
			}
		}	
	}
	return maxIndex, fileList, nil;
}

func (s *StorageManagerImpl) GetNextFileName() string {
	incr := strconv.Itoa(s.index);
	s.index = s.index + 1;
	
	newFileName := s.WorkDir() + "/" + s.Prefix() + incr + s.Suffix();
	
	return newFileName
}

func removeOldFiles(s *StorageManagerImpl) {
	for len(s.fileList) > s.MaxNoOfFiles() {
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
	if file.Size() > s.MinFileSize() {
		s.fileList = append(s.fileList, fileName)
		removeOldFiles(s)
		return nil
	}
	
	return os.Remove(fileName)
}