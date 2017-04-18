package constants


import (
	"os"
)
const (
	COMMAND = "COMMAND" //COMMAND the external program to execute
	WORKDIR = "WORKDIR" //WORKDIR directory to write files to
	TIMEOUT = "TIMEOUT" //TIMEOUT kill the spawned process after set time out
	VIDEOLENGTH = "VIDEOLENGTH" //VIDEOLENGTH the length of the video to record (MILLISECONDS)
	PREFIX = "PREFIX" //PREFIX file name prefix
	SUFFIX = "SUFFIX" //SUFFIX file extension
	MINFILESIZE = "MINFILESIZE" //MINFILESIZE Mininum file size
	MAXNOOFFILES = "MAXNOOFFILES" //MAXNOOFFILES maximum no of files in the WORKDIR
	OPTIONS = "OPTIONS" //OPTIONS command line options
	
	GAP_SERVICE_NAME = "GAPSERVICENAME" // Name of GAP bluetooth service
	GATT_SERVICE_NAME = "GATTSERVICENAME" // Name of GATT bluetooth service
)

const (
	SLASH string = string(os.PathSeparator) //SLASH path separator
)
