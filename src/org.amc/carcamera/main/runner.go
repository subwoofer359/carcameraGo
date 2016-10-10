package main

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	interrupt chan os.Signal
	complete chan error
	timeout <-chan time.Time
	command CameraCommand
}

var ErrTimeout = errors.New("received timeout")

var ErrInterrupt = errors.New("received interrupt")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete: make(chan error),
		timeout: time.After(d),
	}
}

func (r *Runner) add(command CameraCommand) {
	r.command = command
} 

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)
	
	r.complete <- r.command.Run()
	
	select {
		case err :=  <-r.complete:
			return err
		
		case <- r.timeout:
			return ErrTimeout
	}
}


