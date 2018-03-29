package util

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestExecCmd(t *testing.T) {
	fmt.Println(os.Getwd())

	var cmd string
	var err error
	cmd = "ping"
	err = ExecCmd(cmd, "www.baidu.com")
	if err != nil {
		t.Error(err)
		return
	}

	cmd = "ipconfig"
	err = ExecCmd(cmd)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("---------------------")
	time.Sleep(3 * time.Second)
	bys, err := exec.Command("ping", "www.baidu.com").Output()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(bys))
}
