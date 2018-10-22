package runner

import "errors"

//mockStorageManagerForDD Mock StorageManager to work with the
//linux command dd
type mockStorageManagerForDD struct {
	MockStorageManager
}

func (m *mockStorageManagerForDD) GetNextFileName() string {
	return "of=/tmp/e"
}

//GetMockStorageManagerDD returns a mockStorageManagerForDD
// struct for testing with 'dd'
func GetMockStorageManagerDD() *mockStorageManagerForDD {
	return new(mockStorageManagerForDD)
}

//mockStorageManagerForLs Mock StorageManager
// to work with the linux command 'ls'
type mockStorageManagerForLs struct {
	MockStorageManager
}

func (m *mockStorageManagerForLs) GetNextFileName() string {
	return "/tmp"
}

//GetMockStorageManagerLS returns mockStorageManagerForLs
// struct for testing with the linux command 'ls'
func GetMockStorageManagerLS(context map[string]interface{}) *mockStorageManagerForLs {
	t := new(mockStorageManagerForLs)
	t.context = context
	return t
}

//MainMockStorageManager Mock StorageManager
type MainMockStorageManager struct {
	MockStorageManager
}

func (m *MainMockStorageManager) Init() error {
	return errors.New("Test StorageManager init failed")
}

func (m *MainMockStorageManager) GetNextFileName() string {
	return ""
}
