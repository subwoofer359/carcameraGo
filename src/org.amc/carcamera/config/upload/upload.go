// Package upload check a new version of the executable on the mount point
// copy it to the main executable location
// execute the program
package upload

import (
	"io/ioutil"
	"os"

	"errors"

	C "org.amc/carcamera/constants"
)

var (
	filename                         = "main"
	destinationDirectory             = "/home/adrian"
	errNotSameFile                   = errors.New("Not the same file")
	readWriteMask        os.FileMode = 0600
)

func isSignedVersion(context map[string]interface{}) bool {
	//todo not implemented
	return true
}

//checkExecutableExists checks the update executable exists
// and can be deleted afterward. Will return false if not
func checkExecutableExists(context map[string]interface{}) bool {
	fileName := context[C.WORKDIR].(string) + C.SLASH + filename

	fileInfo, err := os.Stat(fileName)

	//File doesnt exist
	if fileInfo == nil || err != nil {
		return false
	}

	readWritePerm := fileInfo.Mode().Perm() & readWriteMask

	return readWritePerm == readWriteMask

}

func copyNewVersionToDirectory(context map[string]interface{}) error {
	fileName := context[C.WORKDIR].(string) + C.SLASH + filename
	if source, err := ioutil.ReadFile(fileName); err != nil {
		return getError(err)
	} else if err := ioutil.WriteFile(destinationDirectory+C.SLASH+filename, source, 0755); err != nil {
		return getError(err)
	}
	return nil
}

func removeOriginalVersion(context map[string]interface{}) error {
	err := os.Remove(context[C.WORKDIR].(string) + C.SLASH + filename)
	return err
}

func getError(err error) error {
	return errors.New("Update Error:" + err.Error())
}

//Verified File on size
//Update to use a hash for checking
func verfiedCopiedFile(context map[string]interface{}) error {
	var (
		sourceFileName = context[C.WORKDIR].(string) + C.SLASH + filename
		destFileName   = destinationDirectory + C.SLASH + filename
		sourceInfo     os.FileInfo
		destInfo       os.FileInfo
		err            error
	)

	if sourceInfo, err = os.Stat(sourceFileName); err != nil {
		return getError(err)
	}

	if destInfo, err = os.Stat(destFileName); err != nil {
		return getError(err)
	}

	if sourceInfo.Size() != destInfo.Size() {
		return getError(errNotSameFile)
	}
	return nil
}

// UpdateExecutable checks and copies new version of executable
// to location for executablr
func UpdateExecutable(context map[string]interface{}) error {
	if checkExecutableExists(context) && isSignedVersion(context) {

		if err := copyNewVersionToDirectory(context); err != nil {
			return getError(err)
		}
		if err := verfiedCopiedFile(context); err != nil {
			return getError(err)
		}
		if err := removeOriginalVersion(context); err != nil {
			return getError(err)
		}
	}
	return nil
}
