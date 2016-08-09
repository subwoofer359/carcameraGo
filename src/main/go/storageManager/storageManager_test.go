package storageManager

import (
    "testing"
    "strconv"
    "io/ioutil"
    "os"
)

var index = 0;

func TestGetNextFileName(t *testing.T) {
	t.Log("test filename creation")
	storage := new(storageManager)
	storage.index = 1;
	
	fileName := storage.getNextFileName()
	t.Log("test First call to getFileName")
	testExpectedFileName(t, fileName)
	
	fileName = storage.getNextFileName()
	t.Log("test second call to getFileName")
	testExpectedFileName(t, fileName)
}

func testExpectedFileName(t *testing.T, fileName string) {
	expectedFileName := getExpectedFileName()
	if(fileName != expectedFileName) {
		t.Errorf("Invalid filename generation: expected(%s), actual(%s)", expectedFileName, fileName);
	}
}

func getExpectedFileName() string {
	index ++;
	return PREFIX + strconv.Itoa(index) + SUFFIX;
}

func TestSetLastIndex(t *testing.T) {
	createTestFile(t)
	expected := 1
	storage := new(storageManager)
	storage.index = 0
	storage.workDir = "/tmp"
	
	storage.setLastIndex()
	if(storage.index != expected) {
		t.Errorf("Last index should return %d but instead returned %d", expected, storage.index);
	}
}

func createTestFile(t *testing.T) {
	info := []byte{};
	err := ioutil.WriteFile("/tmp/video1.mpg", info, os.FileMode(0777))
	if(err == nil) {
		return;
	}
	t.Error(err.Error());
}
