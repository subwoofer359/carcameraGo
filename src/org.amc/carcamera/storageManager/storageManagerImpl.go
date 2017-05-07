package storageManager

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"

	"org.amc/carcamera/check"
	C "org.amc/carcamera/constants"
)

// Filename is left padding with six zeros
var (
	FILENAME_FORMAT      = "%06d"
	FILENAME_INDEX_LIMIT = 999999
	MOUNTED              = true
	REGEXP_NUMBER        = "(\\d+)\\"
)

//StorageManager object
type StorageManagerImpl struct {
	index    int
	fileList []string
	context  map[string]interface{}
}

//New create new StorageManager
func NewStorageManager(context map[string]interface{}) StorageManager {
	s := new(StorageManagerImpl)
	s.context = context
	return s
}

func (s StorageManagerImpl) Prefix() string {
	return s.context[C.PREFIX].(string)
}

func (s StorageManagerImpl) Suffix() string {
	return s.context[C.SUFFIX].(string)
}

func (s StorageManagerImpl) Index() int {
	return s.index
}

func (s StorageManagerImpl) WorkDir() string {
	return s.context[C.WORKDIR].(string)
}

func (s *StorageManagerImpl) SetWorkDir(workDir string) {
	s.context[C.WORKDIR] = workDir
}

func (s StorageManagerImpl) FileList() []string {
	return s.fileList
}

//MaxNoOfFiles return MaxNoOfFiles
func (s StorageManagerImpl) MaxNoOfFiles() int {
	maxfiles, _ := strconv.Atoi(s.context[C.MAXNOOFFILES].(string))
	return maxfiles
}

func (s StorageManagerImpl) MinFileSize() int64 {
	minfileSize, _ := strconv.Atoi(s.context[C.MINFILESIZE].(string))
	return int64(minfileSize)
}

func (s *StorageManagerImpl) GetContext() map[string]interface{} {
	return s.context
}

func (s *StorageManagerImpl) Init() error {
	log.Println("StorageManager Init called")

	if err := check.CheckFileSystem(s.WorkDir(), MOUNTED); err != nil {
		return err
	}

	if index, fileList, err := findAndSaveExistingFileNames(s); err != nil {
		return fmt.Errorf("Error reading Work Directory %s\n", s.WorkDir())
	} else {
		s.index = index
		s.fileList = fileList
	}

	log.Printf("StorageManager: %d previous files found\n", s.index)

	s.index = s.index + 1

	log.Printf("StorageManager:%s\n", s.fileList)

	return nil
}

func findAndSaveExistingFileNames(s *StorageManagerImpl) (int, []string, error) {
	if files, err := ioutil.ReadDir(s.WorkDir()); err != nil {
		return 0, nil, err
	} else {
		maxIndex, fileList := sortFilenames(files, s)
		return maxIndex, fileList, nil
	}

}

func sortFilenames(files []os.FileInfo, s *StorageManagerImpl) (int, []string) {
	const NO_INDEX = 0
	fileList := []string{}
	oldList := []string{}

	pattern := s.Prefix() + REGEXP_NUMBER + s.Suffix()
	maxIndex := 0
	lastIndex := 0
	matcher := regexp.MustCompile(pattern)

	for _, file := range files {
		matches := matcher.FindStringSubmatch(file.Name())
		if foundVideoFile(matches) {
			index := getFileNumber(matches[1])
			fileName := s.WorkDir() + C.SLASH + file.Name()
			if lastIndex == NO_INDEX || isFileNameIndexConsecutive(lastIndex, index) {
				fileList = append(fileList, fileName)
				lastIndex = index
				if index > maxIndex {
					maxIndex = index
				}
			} else {
				oldList = append(oldList, fileName)
			}
		}

	}
	return maxIndex, append(oldList, fileList...)
}

func foundVideoFile(matches []string) bool {
	return len(matches) > 0
}

func getFileNumber(filenumberStr string) int {
	number, _ := strconv.Atoi(filenumberStr)
	return number
}

func isFileNameIndexConsecutive(previousIndex int, index int) bool {
	return index-previousIndex == 1
}

func (s *StorageManagerImpl) GetNextFileName() string {
	if s.index > FILENAME_INDEX_LIMIT {
		s.index = 1
	}

	incr := fmt.Sprintf(FILENAME_FORMAT, s.index)
	s.index = s.index + 1

	newFileName := s.WorkDir() + C.SLASH + s.Prefix() + incr + s.Suffix()

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
		if err := os.Remove(oldFileStr); err != nil {
			log.Println(err)
		}
		s.fileList = s.fileList[1:]
	}
}

func (s *StorageManagerImpl) AddCompleteFile(fileName string) error {

	if file, err := os.Stat(fileName); err != nil {
		return err
	} else if file.Size() >= s.MinFileSize() {
		s.fileList = append(s.fileList, fileName)
		removeOldFiles(s)
		return nil
	} else if s.Index() > 0 {
		s.index = s.index - 1
		log.Println(s.index)
	}

	return os.Remove(fileName)
}
