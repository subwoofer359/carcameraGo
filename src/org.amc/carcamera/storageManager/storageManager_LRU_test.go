package storageManager

import (
	"testing"
)

func TestRemoveLRU(t *testing.T) {
	removeTestFiles()
	
	const NumOfFiles int = 10
	
	t.Log("removing Least recently used")
	
	storage := getNewStorageManager()
	
	for i:= 1; i < NumOfFiles; i++ {
		createTestFile(i, t)
	}
	
	storage.Init();
	
	storage.RemoveLRU()
	
	checkFileDoesntExist(storage.Prefix() + "1" + storage.Suffix(), t)
	
	if len(storage.FileList()) == NumOfFiles {
		t.Error("File should have been removed from file list")
	}
}

func TestRemoveLRUFromEmptyList(t *testing.T) {
	removeTestFiles()
	
	t.Log("removing Least recently used")
	
	storage := getNewStorageManager()
	
	storage.Init();
	
	storage.RemoveLRU()
	
}