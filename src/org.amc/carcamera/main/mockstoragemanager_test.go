package main

import (

)

type mockStorageManagerForDD struct {
	mockStorageManager
}

func (m *mockStorageManagerForDD) GetNextFileName() string {
	return "of=/tmp/e"
}

func GetMockStorageManagerDD() *mockStorageManagerForDD {
	return new(mockStorageManagerForDD)
}

type mockStorageManagerForLs struct {
	mockStorageManager
}

func (m *mockStorageManagerForLs) GetNextFileName() string {
	return "/tmp"
}

func GetMockStorageManagerLS() *mockStorageManagerForLs {
	return new(mockStorageManagerForLs)
}

