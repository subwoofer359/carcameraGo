package storageManager

import (
	C "org.amc/carcamera/constants"
	"io/ioutil"
	"os"
	"log"
	"strings"
	"bytes"
	"testing"
	"fmt"
)

func removeTestFiles() {
	files, err := ioutil.ReadDir(T_WORKDIR);
	
	if err != nil {
		log.Fatal(err);
	}
	
	for _, f := range files {
		if !f.Mode().IsDir() && strings.HasPrefix(f.Name(), T_PREFIX) {
			err = os.Remove(T_WORKDIR + C.SLASH + f.Name());
		}
	}
}

func getNewStorageManager() StorageManager {
	context := map[string] interface{} {
		C.WORKDIR: T_WORKDIR,
		C.TIMEOUT: "5s",
		C.PREFIX: T_PREFIX,
		C.SUFFIX: T_SUFFIX,
		C.MINFILESIZE: "0",
		C.MAXNOOFFILES: "10",
	}
	storage := new(StorageManagerImpl)
	storage.index = 0
	storage.context = context
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
	index := fmt.Sprintf(FILENAME_FORMAT, number)
	err := ioutil.WriteFile(T_WORKDIR + C.SLASH + T_PREFIX + index + T_SUFFIX, *info, os.FileMode(0777))
	if(err == nil) {
		return
	}
	t.Error(err.Error())
}

func checkFileDoesntExist(fileName string, t *testing.T) {
	log.Printf("Checking if file %s exists\n", fileName)
	_, perr := os.Stat(fileName)
	if perr == nil {
		t.Errorf("File %s still exists\n", fileName)
	}
}
