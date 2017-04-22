package storageManager

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	C "org.amc/carcamera/constants"
	"os"
	"testing"
)

var (
	index = 1 //index starting at one
)

const (
	T_WORKDIR = "/tmp"  //T_WORKDIR Directory to use for testing
	T_PREFIX  = "VIDEO" //T_PREFIX filename prefix
	T_SUFFIX  = ".h264" //T_SUFFIX filename extension
)

/*
 * Set up for tests
 */
func TestMain(m *testing.M) {
	MOUNTED = false
	os.Exit(m.Run())
}

func TestGetNextFileName(t *testing.T) {
	removeTestFiles()
	storage := getNewStorageManager()

	fileName := storage.GetNextFileName()
	testExpectedFileName(t, fileName)

	fileName = storage.GetNextFileName()
	testExpectedFileName(t, fileName)
}

func testExpectedFileName(t *testing.T, fileName string) {
	expectedFileName := getExpectedFileName()
	if fileName != expectedFileName {
		t.Errorf("Invalid filename generation: expected(%s), actual(%s)", expectedFileName, fileName)
	}
}

func getExpectedFileName() string {
	if index >= FILENAME_INDEX_LIMIT {
		index = 1
	}
	fileName := T_WORKDIR + C.SLASH + T_PREFIX + fmt.Sprintf(FILENAME_FORMAT, index) + T_SUFFIX
	index++
	return fileName
}

func TestInitForCorrectIndex(t *testing.T) {
	removeTestFiles()

	createTestFile(1, t)
	expected := 2
	storage := getNewStorageManager()

	if storage.Index() != expected {
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

func TestGetNextFileResetsFileLimit(t *testing.T) {
	removeTestFiles()

	FILE_NO_LOWER := 999989
	FILE_NO_UPPER := 999999

	//Create some old files to clean up
	for i := FILE_NO_LOWER; i < FILE_NO_UPPER; i = i + 1 {
		createTestFile(i, t)
	}

	storage := getNewStorageManager()
	index = FILE_NO_UPPER

	storage.GetNextFileName()
	actualFileName := storage.GetNextFileName()
	expectedFileName := getExpectedFileName()
	fmt.Println(expectedFileName)
	if actualFileName != expectedFileName {
		t.Errorf("Filename (%s) doesn't equal %s", actualFileName, expectedFileName)
	}

}

/*
 * Todo: Needs to be fixed
 */
func TestAfterIndexWrapAroundCorrectIndex(t *testing.T) {
	removeTestFiles()
	startID := 999989
	endID := 1000000

	//Create some old files to clean up
	for i := startID; i < endID; i = i + 1 {
		createTestFile(i, t)
	}

	for i := 2; i < 10; i = i + 1 {
		createTestFile(i, t)
	}

	storage := getNewStorageManager()
	log.Println(storage.FileList())

	if storage.Index() != 10 {
		t.Errorf("Filename index (%d) is incorrect", storage.Index())
	}

	fileList := storage.FileList()

	assert.NotNil(t, fileList, "FileList should not be nil")
	assert.NotEmpty(t, fileList, "FileList should not be empty")

	x := 0
	for i := startID; i < endID; i = i + 1 {
		assert.Equal(t, fileList[x], storage.WorkDir()+C.SLASH+T_PREFIX+fmt.Sprintf(FILENAME_FORMAT, i)+T_SUFFIX)
		x = x + 1
	}
}
