package runner

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

type RunnerImpl struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time
	command   CameraCommand
}

var (
	//ErrTimeout error due to process timing out
	ErrTimeout = errors.New("received timeout")

	//ErrInterrupt error due to process bring send a interrupt signal
	ErrInterrupt = errors.New("received interrupt")
)

func (r *RunnerImpl) Add(command CameraCommand) {
	r.command = command
}

func (r *RunnerImpl) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)
	log.Println("Signal Notified")
	go func() {
		r.complete <- r.command.Run()
	}()
	log.Println("command started")

	return r.Handle()
}

func (r *RunnerImpl) Stop() {
	if r.command.Process() != nil {
		r.command.Process().Kill()
	}
}

//Handle return messages from external process
func (r *RunnerImpl) Handle() error {
	select {
	case err := <-r.complete:
		log.Println("command completed")
		return err

	case <-r.timeout:
		log.Println("command Timed out")
		r.Stop()
		return ErrTimeout

	case <-r.interrupt:
		log.Println("command interrupted")
		r.Stop()
		signal.Stop(r.interrupt)
		return ErrInterrupt
	}
}
