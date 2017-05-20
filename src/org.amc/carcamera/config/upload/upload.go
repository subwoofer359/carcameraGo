// Package upload check a new version of the executable on the mount point
// copy it to the main executable location
// execute the program
package upload

import (
	"io/ioutil"
	"log"
	"os"

	"errors"

	C "org.amc/carcamera/constants"
)

var (
	errNotSameFile             = errors.New("Not the same file")
	readWriteMask  os.FileMode = 0600
)

//GetNewUpdater returns initialise Updater struct pointer
func GetNewUpdater(context map[string]interface{}, destinationPath string, filename string) *Updater {
	return &Updater{
		filename,
		destinationPath,
		context,
	}
}

//Updater a data struct for carrying out update operations
type Updater struct {
	filename        string
	destinationPath string
	context         map[string]interface{}
}

func (u Updater) isSignedVersion() bool {
	//todo not implemented
	return true
}

//checkExecutableExists checks the update executable exists
// and can be deleted afterward. Will return false if not
func (u Updater) checkExecutableExists() bool {
	fileName := u.context[C.WORKDIR].(string) + C.SLASH + u.filename

	fileInfo, err := os.Stat(fileName)

	//File doesnt exist
	if fileInfo == nil || err != nil {
		return false
	}

	readWritePerm := fileInfo.Mode().Perm() & readWriteMask

	return readWritePerm == readWriteMask

}

func (u Updater) copyNewVersionToDirectory() error {
	fileName := u.context[C.WORKDIR].(string) + C.SLASH + u.filename
	if source, err := ioutil.ReadFile(fileName); err != nil {
		return getError(err)
	} else if err := ioutil.WriteFile(u.destinationPath+C.SLASH+u.filename, source, 0755); err != nil {
		return getError(err)
	}
	return nil
}

func (u Updater) removeOriginalVersion() error {
	err := os.Remove(u.context[C.WORKDIR].(string) + C.SLASH + u.filename)
	return err
}

func getError(err error) error {
	return errors.New("Update Error:" + err.Error())
}

//Verified File on size
//Update to use a hash for checking
func (u Updater) verfiedCopiedFile() error {
	var (
		sourceFileName = u.context[C.WORKDIR].(string) + C.SLASH + u.filename
		destFileName   = u.destinationPath + C.SLASH + u.filename
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
func (u Updater) UpdateExecutable() error {
	if u.checkExecutableExists() && u.isSignedVersion() {
		log.Printf("Updating file:%s", u.filename)
		if err := u.copyNewVersionToDirectory(); err != nil {
			return getError(err)
		}
		if err := u.verfiedCopiedFile(); err != nil {
			return getError(err)
		}
		if err := u.removeOriginalVersion(); err != nil {
			return getError(err)
		}
		log.Println("Update completed")
	} else {
		log.Println("No Update file found")
	}
	log.Println("No update carried out")
	return nil
}
