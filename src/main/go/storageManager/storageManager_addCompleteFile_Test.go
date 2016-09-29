package storageManager

import (
	"testing"
	"log"
)

func TestAddCompleteFileSizeLessThanAllowed(t *testing.T) {
	removeTestFiles()
	t.Log("removing Least recently used")
	storage := getNewStorageManager()
	
	createEmptyTestFile(0, t)
	
	err := storage.addCompleteFile(storage.GetNextFileName())
	
	if err != nil {
		t.Fatal(err)
	}
	
	checkFileDoesntExist(VIDEO + "0.mpg", t)
	
	checkFileNameNotStored(storage, t)
}

func TestAddCompleteFileStaysWithinFileLimit(t *testing.T) {
	removeTestFiles()
	storage := getNewStorageManager()
	
	for i := 0; i < storage.MaxNoOfFiles() + 10; i = i + 1 {
		createTestFile(i, t)
		storage.addCompleteFile(storage.GetNextFileName())
	}
	
	if len(storage.fileList) != storage.MaxNoOfFiles() {
		t.Error("StorageManager not keeping files created within limits");
	}
}

func checkFileNameNotStored(storage *storageManager, t *testing.T) {
	if len(storage.fileList) > 0 {
		log.Println(storage.fileList)
		t.Error("Invalid file was added to file list")
	}
}