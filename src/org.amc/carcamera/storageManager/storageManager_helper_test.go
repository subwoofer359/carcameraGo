package storageManager

import (
	"io/ioutil"
	"os"
	"log"
	"strings"
	"strconv"
	"bytes"
	"testing"
)

const TMP string = "/tmp"
const VIDEO string = "video"
func removeTestFiles() {
	files, err := ioutil.ReadDir(TMP);
	
	if err != nil {
		log.Fatal(err);
	}
	
	for _, f := range files {
		if !f.Mode().IsDir() && strings.HasPrefix(f.Name(), VIDEO) {
			err = os.Remove(TMP + "/" + f.Name());
		}
	}
}

func getNewStorageManager() *StorageManager {
	storage := new(StorageManager)
	storage.index = 0
	storage.WorkDir = TMP
	storage.Init()
	return storage;
}

func createTestFile(number int, t *testing.T) {
	info, _:= bytes.NewBufferString("Hello World").ReadBytes('d')
	createETestFile(number, &info, t)
}

func createEmptyTestFile(number int, t *testing.T) {
	createETestFile(number, &[]byte{}, t)
}

func createETestFile(number int, info *[]byte, t *testing.T) {
	index := strconv.Itoa(number)
	err := ioutil.WriteFile(TMP + "/" + VIDEO + index + ".mpg", *info, os.FileMode(0777))
	if(err == nil) {
		return
	}
	t.Error(err.Error())
}

func checkFileDoesntExist(fileName string, t *testing.T) {
	_, perr := os.Stat(TMP + "/" + fileName)
	if perr == nil {
		t.Error("Last File still exists")
	}
}
