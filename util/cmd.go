package util

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

func ExecCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	err = cmd.Start()
	if err != nil {
		return err
	}

	reader := bufio.NewReader(stdout)

	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line)
	}
	cmd.Wait()
	return nil
}
