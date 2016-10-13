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
	GetNextFileName() string
	RemoveLRU()
	AddCompleteFile(fileName string) error
}