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

type CameraCommand struct {
	command string
	args []string
	storageManager storageManager.StorageManager
	process *os.Process
	exec func(string, ...string) *exec.Cmd
}

func (c *CameraCommand) Run() error {
	
	filename := c.storageManager.GetNextFileName()
	
	if filename != "" {
		c.args = append(c.args, filename)
	}
	
	cmd := c.exec(c.command, c.args...)
	
	stdout, stderr, pipeError := setOutPipes(cmd)
	
	if pipeError != nil {
		return pipeError
	}
	
	err := cmd.Start()
	
	c.process = cmd.Process
	
	if readErr := readPipe(stdout); readErr != nil {
		return readErr
	}
	
	if readErr := readPipe(stderr); readErr != nil {
		return readErr
	}
	
	if err != nil {
		return errors.New(strings.Join([]string{"ERROR:", err.Error()}, " "))
	} 
	
	cmd.Wait()

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
