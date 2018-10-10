package runner

import (
	"os"
	"time"
)

type SimpleRunnerFactory struct {
}

func (factory *SimpleRunnerFactory) NewRunner(d time.Duration) Runner {
	return &RunnerImpl{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}
