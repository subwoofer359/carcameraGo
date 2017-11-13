package storageManager

import (
	"log"
	"testing"
)

func TestRemoveLRU(t *testing.T) {
	removeTestFiles()

	const NumOfFiles int = 10

	t.Log("removing Least recently used")

	storage := getNewStorageManager()

	for i := 1; i < NumOfFiles; i++ {
		createTestFile(i, t)
	}

	storage.Init()

	storage.RemoveLRU()

	checkFileDoesntExist(storage.Prefix()+"1"+storage.Suffix(), t)

	if len(storage.FileList()) == NumOfFiles {
		t.Error("File should have been removed from file list")
	}
}

func TestRemoveLRUFromEmptyList(t *testing.T) {
	removeTestFiles()

	t.Log("removing Least recently used")

	storage := getNewStorageManager()

	storage.Init()

	storage.RemoveLRU()

}

// TestRemoveOldFile tests if all old files are deleted and reduces the number of files
// to context.MaxNoOfFiles
func TestRemoveOldFile(t *testing.T) {
	removeTestFiles()

	const NumOfFiles int = 20

	t.Log("removing Least recently used")

	storage := getNewStorageManager()

	for i := 1; i <= NumOfFiles; i++ {
		createTestFile(i, t)
	}

	storage.Init()

	log.Printf("Number of files before remove call is %d", len(storage.FileList()))

	if len(storage.FileList()) != NumOfFiles {
		t.Error("Files should have been created")
	}

	storage.RemoveOldFiles()

	log.Printf("Number of files after remove call is %d", len(storage.FileList()))

	if len(storage.FileList()) != storage.MaxNoOfFiles() {
		t.Error("Files should have been removed from file list")
	}
}
