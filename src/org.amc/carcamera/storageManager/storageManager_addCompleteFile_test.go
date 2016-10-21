package storageManager

import (
	"testing"
	"log"
	"strings"
	"math/rand"
	"time"
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
	
	//Create some old files to clean up
	for i := 1; i < 10; i = i + 1 {
		createTestFile(i, t)
	}
	
	storage := getNewStorageManager()
	
	log.Printf("Minimum file size is %d", storage.MinFileSize())
	log.Printf("Maximum no of files is %d", storage.MaxNoOfFiles())
	
	for i := 10; i < storage.MaxNoOfFiles() + 100; i = i + 1 {
		createTestFile(i, t)
		if err := storage.AddCompleteFile(storage.GetNextFileName()); err != nil {
			t.Error(err)
		}
	}
	
	if len(storage.FileList()) > storage.MaxNoOfFiles() {
		t.Errorf("StorageManager not keeping files created within limits.\n" + 
			"Filelimit %d but number of files is %d", storage.MaxNoOfFiles(), len(storage.FileList()));
	}
}

//TestAddCompleteRemovesEmptyFiles creates files greater and less than
// MinFileSize to see if the MaxNoOfFiles is respected
func TestAddCompleteRemovesEmptyFiles(t *testing.T) {
	removeTestFiles()
	
	storage := getNewStorageManager()
	
	log.Printf("Minimum file size is %d", storage.MinFileSize())
	log.Printf("Maximum no of files is %d", storage.MaxNoOfFiles())
	
	
	rand.Seed(time.Now().UnixNano())
	
	for i := 1; i < storage.MaxNoOfFiles() + 100; i = i + 1 {
		r := rand.Intn(10)
		if r < 4 {
			createTestFile(i, t)
		} else {
			createEmptyTestFile(i, t)
		}
		if err := storage.AddCompleteFile(storage.GetNextFileName()); err != nil {
			t.Error(err)
		}
	}
	
	if len(storage.FileList()) > storage.MaxNoOfFiles() {
		t.Errorf("StorageManager not keeping files created within limits.\n" + 
			"Filelimit %d but number of files is %d", storage.MaxNoOfFiles(), len(storage.FileList()));
	}
}

func checkFileNameNotStored(storage StorageManager, t *testing.T) {
	if len(storage.FileList()) > 0 {
		log.Println(storage.FileList())
		t.Error("Invalid file was added to file list")
	}
}