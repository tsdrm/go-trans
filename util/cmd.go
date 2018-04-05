package util

import (
	"errors"
	"os/exec"
	"time"
)

// Default timeout time for shell execution.
var DefaultTimeout = 120 * time.Minute

type Cmder struct {
	*exec.Cmd

	// Timeout for command.
	Timeout time.Duration
}

func NewCmder() *Cmder {
	return &Cmder{
		Timeout: DefaultTimeout,
	}
}

// Set the shell execution timeout.
func (c *Cmder) SetTimeout(timeout time.Duration) {
	c.Timeout = timeout
}

// Start executing the command
//
// name: command name.
// args: command args.
//
// string: standard output string of the command.
// error: error message
func (c *Cmder) Command(name string, args ...string) (string, error) {
	var err error
	var output string
	var bys []byte

	c.Cmd = exec.Command(name, args...)

	var run = make(chan int)

	go func() {
		bys, err = c.Cmd.CombinedOutput()
		run <- 1
	}()
	select {
	case <-run:
		output = string(bys)
	case <-time.After(c.Timeout):
		err = errors.New("time out")
	}

	return output, err
}
