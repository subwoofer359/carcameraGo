package main

import (
	"os"
)

type CameraCommand interface {
	Run() error
	Process() *os.Process
}
