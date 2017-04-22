package check

import (
	"os"
	"os/exec"
	"syscall"
	"errors"
	"regexp"
)

var (
	PERMISSION_ERROR error = errors.New("No permission to file")
	MOUNT_ERROR error = errors.New("Filesystem not mounted")
)

func CheckFileSystem(path string, mounted bool) error {
	if err := checkPathExist(path); err != nil {
		return err
	}
	
//	if err := checkIsOwnerOfPath(path); err != nil {
//		return err
//	}
	
	if mounted {
		if err := checkIsMounted(path); err != nil {
			return err
		}
	}
	if err := isDirectoryWritable(path); err != nil {
		return err
	}
	return nil
}

func checkPathExist(path string) error {
	if file, err := os.Open(path); os.IsNotExist(err) {
		return err
	} else {
		file.Close()
		return nil
	}
}

func checkIsOwnerOfPath(path string) error {
	fileInfo, err := os.Stat(path) 
	if err != nil {
		return err
	}
	
	if fileInfo.Sys().(*syscall.Stat_t).Uid != uint32(os.Getuid()) {
		return PERMISSION_ERROR
	} else {
		return nil
	}
}

func checkIsMounted(path string) (error) {
	cmd := exec.Command("mount")
	
	if result, err := cmd.Output(); err != nil {
		return err
	} else {
		
		output := string(result)
		pattern := path +"\\s"
		matcher := regexp.MustCompile(pattern)
		
		if matcher.FindStringIndex(output) == nil {
			return MOUNT_ERROR
		} else {
			return nil
		}
	}
}

// isDirectoryWritable
// param directory string directory to test for writing to
func isDirectoryWritable(path string) error {
	const (
		TESTFILENAME string = "test"
		SLASH string = string(os.PathSeparator)
	)
	
	if file, err := os.Create(path + SLASH + TESTFILENAME); err != nil {
		return err
	} else {
		_, err := file.WriteString("Test")
		return err
	}
	
}