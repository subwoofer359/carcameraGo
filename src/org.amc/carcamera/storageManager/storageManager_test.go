package storageManager

import (
	C "org.amc/carcamera/constants"
    "testing"
    "strconv"
)

var (
		index = 1 //index starting at one
)
const (
		T_WORKDIR = "/tmp" //T_WORKDIR Directory to use for testing
		T_PREFIX = "VIDEO" //T_PREFIX filename prefix
		T_SUFFIX = ".h264" //T_SUFFIX filename extension
)

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
	fileName := T_WORKDIR + C.SLASH + T_PREFIX + strconv.Itoa(index) + T_SUFFIX
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
	storage := getNewStorageManager()
	storage.SetWorkDir("/root")
	err := storage.Init()
	if err == nil {
		t.Error("Should have thrown an error on attempt to read directory")
	}
}






