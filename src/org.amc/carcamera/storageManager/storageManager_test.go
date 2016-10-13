package storageManager

import (
    "testing"
    "strconv"
)

var index = 1

func TestGetNextFileName(t *testing.T) {
	removeTestFiles();
	storage := getNewStorageManager()	
	
	fileName := storage.GetNextFileName()
	testExpectedFileName(t, fileName)
	
	fileName = storage.GetNextFileName()
	testExpectedFileName(t, fileName)
}

func testExpectedFileName(t *testing.T, fileName string) {
	expectedFileName := getExpectedFileName()
	if(fileName != expectedFileName) {
		t.Errorf("Invalid filename generation: expected(%s), actual(%s)", expectedFileName, fileName)
	}
}

func getExpectedFileName() string {
	fileName :=TMP + "/" + PREFIX + strconv.Itoa(index) + SUFFIX
	index ++
	return fileName;
}

func TestInitForCorrectIndex(t *testing.T) {
	removeTestFiles()
	
	createTestFile(1, t)
	expected := 2
	storage := getNewStorageManager()
	
	if(storage.Index() != expected) {
		t.Errorf("Last index should return %d but instead returned %d", expected, storage.Index())
	}
}

func TestInitCantReadWorkDir(t *testing.T) {
	removeTestFiles()
	storage := new(StorageManagerImpl)
	storage.index = 0
	storage.workDir = "/root"
	err := storage.Init()
	if err == nil {
		t.Error("Should have thrown an error on attempt to read directory")
	}
}






