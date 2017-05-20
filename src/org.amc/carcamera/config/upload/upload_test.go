package upload

import (
	"bytes"
	"log"
	"testing"

	"io/ioutil"
	"os"

	"fmt"

	"github.com/stretchr/testify/assert"
	C "org.amc/carcamera/constants"
)

var (
	context            map[string]interface{}
	testDir            = "/tmp"
	testFileName       = "Alf"
	testDestDirectory  = "/tmp/copied"
	testCopiedFilePath = testDestDirectory + C.SLASH + testFileName
	testFilePath       = testDir + C.SLASH + testFileName
)

func TestMain(m *testing.M) {
	setUp()
	os.Exit(m.Run())
}

func setUp() {
	context = make(map[string]interface{})

	//Replace defaults to test values
	context[C.WORKDIR] = testDir

	filename = testFileName
	destinationDirectory = testDestDirectory
}

func TestCheckExecutableExists(t *testing.T) {
	context[C.WORKDIR] = "/bin"
	filename = "ls"

	exists := checkExecutableExists(context)
	assert.True(t, exists, "File should exist")
}

func TestCheckExecutableNotExists(t *testing.T) {
	context[C.WORKDIR] = "/bin"
	filename = "Alf"

	exists := checkExecutableExists(context)
	assert.False(t, exists, "File should not exist")
}

func TestCheckExecutableExistsNotDeletable(t *testing.T) {
	setUp()
	setUpFileTest()
	removePermissionsFromFile(t, testFilePath)
	log.Println(testFilePath)

	exists := checkExecutableExists(context)

	assert.False(t, exists, "File exists but can not be deleted")

	restorePermissionsFromFile(t, testFilePath)
}

func createCopiedDir(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.Mkdir(testDestDirectory, 0755)
	} else {
		log.Println("Directory already exists:" + path)
	}
}

func removeTestDirectory(path string) {
	os.RemoveAll(path)
}

func createTestFile(path string) {
	data := bytes.NewBufferString("Test File").Bytes()
	if err := ioutil.WriteFile(path, data, 0775); err != nil {
		log.Println(err)
	}
}

func setUpFileTest() {
	removeTestDirectory(testDestDirectory)
	createCopiedDir(testDestDirectory)
	createTestFile(testFilePath)
}

func TestCopyNewVersionToDirectory(t *testing.T) {
	setUp()
	setUpFileTest()

	err := copyNewVersionToDirectory(context)

	assert.Nil(t, err, "Should be no error on copy")

	_, err = os.Stat(testCopiedFilePath)

	assert.Nil(t, err, fmt.Sprintf("Should exist:%s", testCopiedFilePath))
}

func TestCopyNewVersionToDirectoryCantRead(t *testing.T) {
	setUp()
	setUpFileTest()

	removePermissionsFromFile(t, testFilePath)

	err := copyNewVersionToDirectory(context)

	assert.NotNil(t, err, "Should be error on copy")

	restorePermissionsFromFile(t, testFilePath)
}

func TestCopyNewVersionToDirectoryCantWrite(t *testing.T) {
	setUp()
	setUpFileTest()

	removePermissionsFromFile(t, testDestDirectory)

	err := copyNewVersionToDirectory(context)

	assert.NotNil(t, err, "Should be error on copy")

	restorePermissionsFromFile(t, testDestDirectory)
}

func removePermissionsFromFile(t *testing.T, path string) {
	if err := os.Chmod(path, 0000); err != nil {
		t.Error(err)
	}
}

func restorePermissionsFromFile(t *testing.T, path string) {
	if err := os.Chmod(path, 0755); err != nil {
		t.Error(err)
	}
}

func TestVerfiedCopiedFile(t *testing.T) {
	setUp()
	setUpFileTest()
	copyNewVersionToDirectory(context)

	assert.Nil(t, verfiedCopiedFile(context), "Files should be the same")
}

func TestVerfiedCopiedFileNotSameAsOriginal(t *testing.T) {
	setUp()
	setUpFileTest()

	data := bytes.NewBufferString("Test File Error").Bytes()
	if err := ioutil.WriteFile(testCopiedFilePath, data, 0775); err != nil {
		log.Println(err)
	}

	assert.NotNil(t, verfiedCopiedFile(context), "Files are not the same")
}

func TestRemoveOriginalFile(t *testing.T) {
	setUp()
	setUpFileTest()

	if _, err := os.Stat(testFilePath); err != nil {
		assert.Nil(t, err, "File should exist before deletion")
	}

	err := removeOriginalVersion(context)

	if _, err := os.Stat(testFilePath); err != nil {
		assert.NotNil(t, err, "File not should exist before deletion")
	}

	assert.Nil(t, err, "Should be no error on deletion")
}
