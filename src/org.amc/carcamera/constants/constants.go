package constants


import (
	"os"
)
const (
	WORKDIR = iota //WORKDIR directory to write files to
	TIMEOUT = iota //TIMEOUT kill the spawned process after set time out
	PREFIX = iota //PREFIX file name prefix
	SUFFIX = iota //SUFFIX file extension
	MINFILESIZE = iota //MINFILESIZE Mininum file size
	MAXNOOFFILES = iota //MAXNOOFFILES maximum no of files in the WORKDIR
)

const (
	SLASH string = string(os.PathSeparator)
)

