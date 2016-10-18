package main

type mockStorageManager struct {
	index int
	workDir string
	context map[string] string
}

func (m mockStorageManager) Init() error {
	return nil
}

func (m mockStorageManager) Prefix() string {
	return "video"
}

func (m mockStorageManager) Suffix() string {
	return ".mpg"
}

func (m mockStorageManager) Index() int {
	return m.index
}

func (m mockStorageManager) WorkDir() string {
	return m.workDir
}

func (m *mockStorageManager) SetWorkDir(workDir string) {
	m.workDir = workDir
}

func (m mockStorageManager) FileList() []string {
	return []string{}
}

func (m mockStorageManager) MaxNoOfFiles() int {
	return 10
}

func (m mockStorageManager) MinFileSize() int64 {
	return 0
}

func (m mockStorageManager) RemoveLRU() {
	
}

func (m mockStorageManager) AddCompleteFile(fileName string) error {
	return nil
}

func (m *mockStorageManager) GetContext() map[string] string {
	return m.context
}
