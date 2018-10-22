package runner

import C "org.amc/carcamera/constants"

type MockStorageManager struct {
	index   int
	context map[string]interface{}
}

func (m MockStorageManager) Init() error {
	return nil
}

func (m MockStorageManager) Prefix() string {
	return "video"
}

func (m MockStorageManager) Suffix() string {
	return ".mpg"
}

func (m MockStorageManager) Index() int {
	return m.index
}

func (m MockStorageManager) WorkDir() string {
	return m.context[C.WORKDIR].(string)
}

func (m *MockStorageManager) SetWorkDir(workDir string) {
	m.context[C.WORKDIR] = workDir
}

func (m MockStorageManager) FileList() []string {
	return []string{}
}

func (m MockStorageManager) MaxNoOfFiles() int {
	return 10
}

func (m MockStorageManager) MinFileSize() int64 {
	return 0
}

func (m MockStorageManager) RemoveLRU() {

}

func (m MockStorageManager) RemoveOldFiles() {

}

func (m MockStorageManager) AddCompleteFile(fileName string) error {
	return nil
}

func (m *MockStorageManager) GetContext() map[string]interface{} {
	return m.context
}
