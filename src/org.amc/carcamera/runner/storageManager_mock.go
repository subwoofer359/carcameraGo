package runner

type MockStorageManager struct {
	index   int
	workDir string
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
	return m.workDir
}

func (m *MockStorageManager) SetWorkDir(workDir string) {
	m.workDir = workDir
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
