package main

import (
	"os"
	"os/exec"
	"errors"
	"strings"
	"log"
	"io"
	"bufio"
	"org.amc/carcamera/storageManager"
)

type CameraCommandImpl struct {
	command string
	args []string
	storageManager storageManager.StorageManager
	process *os.Process
	exec func(string, ...string) *exec.Cmd
}

func (c CameraCommandImpl) Process() *os.Process {
	return c.process
}

func (c *CameraCommandImpl) Run() error {
		
	filename := c.storageManager.GetNextFileName()
	
	cmd := c.exec(c.command, append(c.args, filename)...)
	
	stdout, stderr, pipeError := setOutPipes(cmd)
	
	if pipeError != nil {
		return pipeError
	}
	
	if err := cmd.Start(); err != nil {
		return errors.New(strings.Join([]string{"ERROR:", err.Error()}, " "))
	}
	
	c.process = cmd.Process
	
	if readErr := readPipe(stdout); readErr != nil {
		return readErr
	}
	
	if readErr := readPipe(stderr); readErr != nil {
		return readErr
	} 
	
	if err := cmd.Wait(); err != nil {
		return err
	}

	c.storageManager.AddCompleteFile(filename)
	return errors.New("completed")
}

type commandObject interface {
	StdoutPipe() (io.ReadCloser, error)
	StderrPipe() (io.ReadCloser, error)
}

func setOutPipes(cmd commandObject) (io.Reader, io.Reader, error) {
	stdout, err := cmd.StdoutPipe()
	
	if err != nil {
		return nil, nil, errors.New("Can't connect to STDOUT")	
	}
	stderr, err := cmd.StderrPipe()
	
	if err != nil {
		return nil, nil, errors.New("Can't connect to STDERR")
	}
		
	return stdout, stderr, nil
}

func readPipe(pipe io.Reader) error {
	in := bufio.NewScanner(pipe)
	
	for in.Scan() {
		log.Printf(in.Text())
	}
	
	return in.Err()
}
