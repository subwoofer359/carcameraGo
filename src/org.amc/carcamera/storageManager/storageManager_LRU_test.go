package storageManager

import (
	"testing"
)

func TestRemoveLRU(t *testing.T) {
	removeTestFiles()
	
	const NumOfFiles int = 10
	
	t.Log("removing Least recently used")
	
	storage := getNewStorageManager()
	
	for i:= 0; i < NumOfFiles; i++ {
		createTestFile(i, t)
	}
	
	storage.Init();
	
	storage.RemoveLRU()
	
	checkFileDoesntExist(VIDEO + "0.mpg", t)
	
	if len(storage.fileList) == NumOfFiles {
		t.Error("File should have been removed from file list")
	}
}