package storageManager

import (
	"testing"
	"log"
	"strings"
)

func TestAddCompleteFileSizeLessThanAllowed(t *testing.T) {
	removeTestFiles()
	t.Log("removing Least recently used")
	storage := getNewStorageManager()
	
	createEmptyTestFile(1, t)
	
	filename := storage.GetNextFileName()
	
	err := storage.AddCompleteFile(filename)
	
	if err != nil {
		t.Fatal(err)
	}
	
	checkFileDoesntExist(storage.Prefix() + "1" + storage.Suffix(), t)
	
	checkFileNameNotStored(storage, t)
}

func TestAddCompleteFileForNonExistingFile(t *testing.T) {
	removeTestFiles()
	t.Log("removing Least recently used")
	storage := getNewStorageManager()
	
	filename := storage.GetNextFileName()
	
	err := storage.AddCompleteFile(filename)
	
	if err == nil || !strings.Contains(err.Error(), "no such file or directory") {
		t.Fatal("Should have thrown 'No such file' exception")
	}
}

func TestAddCompleteFileStaysWithinFileLimit(t *testing.T) {
	removeTestFiles()
	storage := getNewStorageManager()
	
	for i := 1; i < storage.MaxNoOfFiles() + 10; i = i + 1 {
		createTestFile(i, t)
		storage.AddCompleteFile(storage.GetNextFileName())
	}
	
	if len(storage.FileList()) != storage.MaxNoOfFiles() {
		t.Error("StorageManager not keeping files created within limits");
	}
}

func checkFileNameNotStored(storage StorageManager, t *testing.T) {
	if len(storage.FileList()) > 0 {
		log.Println(storage.FileList())
		t.Error("Invalid file was added to file list")
	}
}