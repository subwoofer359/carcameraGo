package storageManager

import (
	C "org.amc/carcamera/constants"
	"os"
	"testing"
)


func TestDestinationIsWritable (t *testing.T) {
	removeTestFiles();
	const TESTDIR string = T_WORKDIR + "/testdir" 
	
	if err := os.RemoveAll(TESTDIR); err != nil {
		t.Error(err)
	}
	
	if err := os.Mkdir(TESTDIR, 0550); err != nil {
		t.Error(err)
	}
	
	
	
	context := map[string] string {
		C.WORKDIR: TESTDIR,
		C.TIMEOUT: "5s",
		C.PREFIX: T_PREFIX,
		C.SUFFIX: T_SUFFIX,
		C.MINFILESIZE: "0",
		C.MAXNOOFFILES: "10",
	}
	
	storage := new(StorageManagerImpl)
	storage.index = 0
	storage.context = context
	if err := storage.Init(); err == nil {
		t.Error("Error should be thrown as storageManager can't write to directory")
	}
}
