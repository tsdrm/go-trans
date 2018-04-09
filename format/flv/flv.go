package flv

import (
	"bufio"
	"github.com/tangs-drm/go-trans"
	"github.com/tangs-drm/go-trans/log"
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

	// Get file info from ffprobe
	var cmder = util.NewCmder()
	var cmdOutput, err = cmder.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", input)
	if err != nil {
		log.E("TYPE_FLV command ffprobe with input: %v error: %v", input, err)
		return go_trans.TransCommandError, go_trans.TransMessage{}, err
	}
	var cmdFormat = util.Map{}
	err = util.Json2S(cmdOutput, &cmdFormat)
	if err != nil {
		log.E("TYPE_FLV trans file format: %v to Map with input: %v error: %v", cmdOutput, input, err)
		return go_trans.TransCommandError, go_trans.TransMessage{}, err
	}
	var format = cmdFormat.Map("format")

	var beg, end int64 = util.Now13(), 0
	// Start to exec transcoding task.
	flv.Cmd = exec.Command("ffmpeg", params...)
	flv.Cmd.Stdout = flv.Cmd.Stderr
	stdout, err := flv.Cmd.StdoutPipe()
	if err != nil {
		log.E("TYPE_FLV command get StdoutPipe with input: %v, output: %v error: %v", input, output, err)
		return go_trans.TransSystemError, go_trans.TransMessage{}, err
	}

	err = flv.Cmd.Start()
	if err != nil {
		log.E("TYPE_FLV command start with input: %v, output: %v error: %v", input, output, err)
		return go_trans.TransCommandError, go_trans.TransMessage{}, err
	}

	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if io.EOF != err {
				log.E("TYPE_FLV reader read string with input: %v, output: %v error: %v", input, output, err)
				return go_trans.TransCommandError, go_trans.TransMessage{}, err
			}
			break
		}
		flv.Lines = append(flv.Lines, line)
	}
	err = flv.Cmd.Wait()
	if err != nil {
		log.E("TYPE_FLV command wait with input: %v, output: %v error: %v", input, output, err)
		return go_trans.TransCommandError, go_trans.TransMessage{}, err
	}
	end = util.Now13()

	var message = go_trans.TransMessage{
		Input: go_trans.InputFile{
			Cdn: "default",
		},
		Size:         format.Int("size"),
		Position:     "",
		Cost:         int(end - beg),
		CreationTime: beg,
		Duration:     format.Float32("duration"),
	}

	return go_trans.TransOk, message, nil
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
