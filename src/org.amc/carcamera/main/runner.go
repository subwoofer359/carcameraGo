package main

import (
	"errors"
	"os"
	"os/signal"
	"time"
	"log"
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
	log.Println("Signal Notified")
	go func() {
		r.complete <- r.command.Run()
	}()
	log.Println("command started")
	
	select {
		case err :=  <-r.complete:
			log.Println("command completed")
			return err
		
		case <- r.timeout:
			log.Println("command Timed out")
			r.stop()
			return ErrTimeout
		
		case <-r.interrupt:
			log.Println("command interrupted")
			r.stop() 
			signal.Stop(r.interrupt)
			return ErrInterrupt
	}
}

func (r *Runner) stop() {
	if r.command.Process() != nil {
		r.command.Process().Kill()
	}
}


