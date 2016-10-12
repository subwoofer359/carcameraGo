package main

import (
	 "org.amc/carcamera/storageManager"
)
type mockStorageManager struct {
	index int
	workDir string
}

func (m mockStorageManager) Init() {
	
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

func (m mockStorageManager) GetNextFileName() string {
	return ""
}

func (m mockStorageManager) RemoveLRU() {
	
}

func (m mockStorageManager) AddCompleteFile(fileName string) error {
	return nil
}

func GetMockStorageManager() storageManager.StorageManager {
	return new(mockStorageManager)
}
