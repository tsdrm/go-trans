package flv

import (
	"bufio"
	"github.com/tangs-drm/go-trans"
	"github.com/tangs-drm/go-trans/util"
	"io"
	"os/exec"
)

type Flv struct {
	Cmd   *exec.Cmd
	Lines []string
}

const TYPE_FLV string = "flv"

func (flv *Flv) Type() string {
	return TYPE_FLV
}

func (flv *Flv) Exec(input, output string, args map[string]string) (int, go_trans.TransMessage, error) {
	var params = []string{}
	for k, v := range args {
		params = append(params, k, v)
	}

	flv.Cmd = exec.Command("ffmpeg", params...)
	flv.Cmd.Stdout = flv.Cmd.Stderr
	stdout, err := flv.Cmd.StdoutPipe()
	if err != nil {
		return go_trans.TransSystemError, go_trans.TransMessage{}, err
	}

	err = flv.Cmd.Start()
	if err != nil {
		return go_trans.TransCommandError, go_trans.TransMessage{}, err
	}

	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		flv.Lines = append(flv.Lines, line)
	}
	flv.Cmd.Wait()

	if err != nil {
		return go_trans.TransCommandError, go_trans.TransMessage{}, err
	}

	return go_trans.TransOk, go_trans.TransMessage{}, nil
}

func (flv *Flv) Pid() int {
	if flv.Cmd == nil {
		return -1
	}
	return flv.Cmd.Process.Pid
}

func (flv *Flv) Cancel() error {
	if flv.Cmd == nil {
		return util.NewError("%v", go_trans.TransProcessNotExist)
	}

	return flv.Cmd.Process.Kill()
}

func (flv *Flv) Progress() (map[string]interface{}, error) {
	return nil, nil
}
