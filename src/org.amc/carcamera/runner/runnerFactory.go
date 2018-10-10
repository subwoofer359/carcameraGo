package runner

import (
	"time"
)

//Runner creates a runner object
type RunnerFactory interface {
	NewRunner(d time.Duration) Runner
}
