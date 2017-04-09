package check

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "os"
)

var path string = "/tmp"
const TESTDIR string = "/tmp/testdir" 


func TestCheckFSDirectoryDoesntExist(t *testing.T) {
	wrongPath := "/tmpf"
	
	err := checkPathExist(wrongPath)
	
	if err == nil {
		t.Error("Should be an error for non existing directory")
		t.Fatal(err)
	}
	
	assert.NotNil(t, err, "FileSystem Check should not pass")
}

func TestCheckFSDirectoryIsNotOwner(t *testing.T) {
	wrongPath := "/bin"
	
	err := checkIsOwnerOfPath(wrongPath)
	
	if err == nil {
		t.Error("Should be an error for not owner")
		t.Fatal(err)
	}
	
	assert.NotNil(t, err, "FileSystem Check should not pass")
}

func TestCheckFSDirectoryShouldBeNotMount(t *testing.T) {
	wrongPath := "/bin"
	
	err := checkIsMounted(wrongPath)
	
	if err == nil {
		t.Error("Should be an error for not being mount")
		t.Fatal(err)
	}
	
	assert.NotNil(t, err, "FileSystem Check should not pass")
}

func TestCheckFSDirectoryShouldBeMount(t *testing.T) {
	wrongPath := "/sys/fs/cgroup"
	
	err := checkIsMounted(wrongPath)
	
	if err != nil {
		t.Error("Should be an no error for not being mount")
		t.Fatal(err)
	}
	
	assert.Nil(t, err, "FileSystem Check should pass")
}

func TestDestinationIsWritable (t *testing.T) {
	const TESTDIR string = "/tmp/testdir" 
	
	setupTestDir(t, 0777)
	
	writeErr := isDirectoryWritable(TESTDIR)
	
	assert.Nil(t, writeErr, "FileSystem should be writable")
	
}

func setupTestDir(t *testing.T, perm os.FileMode) {
	if err := os.RemoveAll(TESTDIR); err != nil {
		t.Error(err)
	}
	
	if err := os.Mkdir(TESTDIR, perm); err != nil {
		t.Error(err)
	}
}

func TestDestinationIsNotWritable (t *testing.T) {
	
	setupTestDir(t, 0550)
	
	writeErr := isDirectoryWritable(TESTDIR)
	
	assert.NotNil(t, writeErr, "FileSystem should NOT be writable")
	
}

func TestCheckFileSystem(t *testing.T) {
	setupTestDir(t, 0777)
	err := CheckFileSystem(TESTDIR, false)
	
	assert.Nil(t, err, "Test should pass")
}