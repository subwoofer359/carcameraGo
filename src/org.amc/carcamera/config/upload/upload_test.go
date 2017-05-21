package upload

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"os"

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

	updater Updater
)

func TestMain(m *testing.M) {
	setUp()
	os.Exit(m.Run())
}

func setUp() {
	context = make(map[string]interface{})

	//Replace defaults to test values
	context[C.WORKDIR] = testDir

	updater = Updater{
		filename:        testFileName,
		destinationPath: testDestDirectory,
		context:         context,
	}
}

func TestCheckExecutableExists(t *testing.T) {
	context[C.WORKDIR] = "/bin"

	updater = Updater{
		"ls",
		testDir,
		context,
	}

	exists := updater.checkExecutableExists()
	assert.True(t, exists, "File should exist")
}

func TestCheckExecutableNotExists(t *testing.T) {
	context[C.WORKDIR] = "/bin"

	updater = Updater{
		"Alf",
		testDir,
		context,
	}

	exists := updater.checkExecutableExists()
	assert.False(t, exists, "File should not exist")
}

func TestCheckExecutableExistsNotDeletable(t *testing.T) {
	setUp()
	setUpFileTest()
	removePermissionsFromFile(t, testFilePath)

	exists := updater.checkExecutableExists()

	assert.False(t, exists, "File exists but can not be deleted")

	restorePermissionsFromFile(t, testFilePath)
}

func setUpFileTest() {
	removeTestDirectory(testDestDirectory)
	createCopiedDir(testDestDirectory)
	createTestFile(testFilePath)
}

func removeTestDirectory(path string) {
	os.RemoveAll(path)
}

func createCopiedDir(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.Mkdir(testDestDirectory, 0755)
	} else {
		log.Println("Directory already exists:" + path)
	}
}

func createTestFile(path string) {
	data := bytes.NewBufferString("Test File").Bytes()
	if err := ioutil.WriteFile(path, data, 0775); err != nil {
		log.Println(err)
	}
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

func TestCopyNewVersionToDirectory(t *testing.T) {
	setUp()
	setUpFileTest()

	err := updater.copyNewVersionToDirectory()

	assert.Nil(t, err, "Should be no error on copy")

	_, err = os.Stat(testCopiedFilePath)

	assert.Nil(t, err, fmt.Sprintf("Should exist:%s", testCopiedFilePath))
}

func TestCopyNewVersionToDirectoryCantRead(t *testing.T) {
	setUp()
	setUpFileTest()

	removePermissionsFromFile(t, testFilePath)

	err := updater.copyNewVersionToDirectory()

	assert.NotNil(t, err, "Should be error on copy")

	restorePermissionsFromFile(t, testFilePath)
}

func TestCopyNewVersionToDirectoryCantWrite(t *testing.T) {
	setUp()
	setUpFileTest()

	removePermissionsFromFile(t, testDestDirectory)

	err := updater.copyNewVersionToDirectory()

	assert.NotNil(t, err, "Should be error on copy")

	restorePermissionsFromFile(t, testDestDirectory)
}

func TestVerfiedCopiedFile(t *testing.T) {
	setUp()
	setUpFileTest()
	updater.copyNewVersionToDirectory()

	assert.Nil(t, updater.verfiedCopiedFile(), "Files should be the same")
}

func TestVerfiedCopiedFileNotSameAsOriginal(t *testing.T) {
	setUp()
	setUpFileTest()

	data := bytes.NewBufferString("Test File Error").Bytes()
	if err := ioutil.WriteFile(testCopiedFilePath, data, 0775); err != nil {
		log.Println(err)
	}

	assert.NotNil(t, updater.verfiedCopiedFile(), "Files are not the same")
}

func TestRemoveOriginalFile(t *testing.T) {
	setUp()
	setUpFileTest()

	if _, err := os.Stat(testFilePath); err != nil {
		assert.Nil(t, err, "File should exist before deletion")
	}

	err := updater.removeOriginalVersion()

	if _, err := os.Stat(testFilePath); err != nil {
		assert.NotNil(t, err, "File not should exist before deletion")
	}

	assert.Nil(t, err, "Should be no error on deletion")
}

func TestUpdateExecutable(t *testing.T) {
	setUp()
	setUpFileTest()
	err := updater.UpdateExecutable()

	assert.NoError(t, err, "Should be no error")
}

func TestUpdateNoExecutable(t *testing.T) {

	updater.filename = "nofile"

	err := updater.UpdateExecutable()

	assert.NoError(t, err, "Should be no error")
}

func TestUpdateExecutableError(t *testing.T) {
	setUp()
	setUpFileTest()

	//Directory with no permission to write to
	removePermissionsFromFile(t, testDestDirectory)
	err := updater.UpdateExecutable()

	restorePermissionsFromFile(t, testDestDirectory)
	assert.Error(t, err, "Should be an error thrown")
	log.Println(err)
}

func TestNewUpdater(t *testing.T) {
	u := NewUpdater(context, testDestDirectory, testFileName)
	assert.Equal(t, testDestDirectory, u.destinationPath, "Paths should be the same")
	assert.Equal(t, testFileName, u.filename, "Filenames should be the same")
}
