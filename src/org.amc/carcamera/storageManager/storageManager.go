package storageManager

type StorageManager interface {
	Init() error
	Prefix() string
	Suffix() string
	Index() int
	WorkDir() string
	SetWorkDir(workDir string)
	FileList() []string
	MaxNoOfFiles() int
	MinFileSize() int64
	GetNextFileName() string
	RemoveLRU()
	RemoveOldFiles()
	AddCompleteFile(fileName string) error
	GetContext() map[string]interface{}
}
