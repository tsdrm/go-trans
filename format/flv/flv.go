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

// it not used now.
func (flv *Flv) _Exec(input, output string, args util.Map) (int, go_trans.TransMessage, go_trans.Error) {
	var params = []string{}
	params = append(params, "-i", input)
	for k := range args {
		params = append(params, k, args.String(k))
	}
	params = append(params, "-y", output)

	// Get file info from ffprobe
	var cmder = util.NewCmder()
	var cmdOutput, err = cmder.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", input)
	if err != nil {
		log.E("TYPE_FLV command ffprobe with input: %v error: %v", input, err)
		return go_trans.TransCommandError, go_trans.TransMessage{}, go_trans.Error{Err: err}
	}
	var cmdFormat = util.Map{}
	err = util.Json2S(cmdOutput, &cmdFormat)
	if err != nil {
		log.E("TYPE_FLV trans file format: %v to Map with input: %v error: %v", cmdOutput, input, err)
		return go_trans.TransCommandError, go_trans.TransMessage{}, go_trans.Error{Err: err}
	}
	var format = cmdFormat.Map("format")

	// Start to exec transcoding task.
	var beg, end int64 = util.Now13(), 0
	flv.Cmd = exec.Command("ffmpeg", params...)
	//flv.Cmd.Stdout = flv.Cmd.Stderr
	stderr, err := flv.Cmd.StderrPipe()
	if err != nil {
		log.E("TYPE_FLV command get StdoutPipe with input: %v, output: %v, params: %v error: %v", input, output, params, err)
		return go_trans.TransSystemError, go_trans.TransMessage{}, go_trans.Error{Err: err}
	}

	err = flv.Cmd.Start()
	if err != nil {
		log.E("TYPE_FLV command start with input: %v, output: %v, params: %v error: %v", input, output, params, err)
		return go_trans.TransCommandError, go_trans.TransMessage{}, go_trans.Error{Err: err}
	}

	reader := bufio.NewReader(stderr)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if io.EOF != err {
				log.E("TYPE_FLV reader read string with input: %v, output: %v, params: %v error: %v", input, output, params, err)
				return go_trans.TransCommandError, go_trans.TransMessage{}, go_trans.Error{Err: err, Lines: flv.Lines}
			}
			break
		}
		flv.Lines = append(flv.Lines, line)
	}
	err = flv.Cmd.Wait()
	if err != nil {
		log.E("TYPE_FLV command wait with input: %v, output: %v, params: %v error: %v", input, output, params, err)
		return go_trans.TransCommandError, go_trans.TransMessage{}, go_trans.Error{Err: err}
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
	log.D("TYPE_FLV command with input: %v, output: %v, params: %v success with message: %v", input, output, params, util.S2Json(message))

	return go_trans.TransOk, message, go_trans.Error{}
}

func (flv *Flv) Exec(input, output string, args util.Map) (int, go_trans.TransMessage, go_trans.Error) {
	var params = []string{}
	params = append(params, "-i", input)
	for k := range args {
		params = append(params, k, args.String(k))
	}
	params = append(params, "-y", output)

	// Get file info from ffprobe
	var cmder = util.NewCmder()
	var cmdOutput, err = cmder.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", input)
	if err != nil {
		log.E("TYPE_FLV command ffprobe with input: %v error: %v", input, err)
		return go_trans.TransCommandError, go_trans.TransMessage{}, go_trans.Error{Err: err}
	}
	var cmdFormat = util.Map{}
	err = util.Json2S(cmdOutput, &cmdFormat)
	if err != nil {
		log.E("TYPE_FLV trans file format: %v to Map with input: %v error: %v", cmdOutput, input, err)
		return go_trans.TransCommandError, go_trans.TransMessage{}, go_trans.Error{Err: err}
	}
	var format = cmdFormat.Map("format")

	// Start to exec transcoding task.
	var beg, end int64 = util.Now13(), util.Now13()
	cmder = util.NewCmder()
	cmdOutput, err = cmder.Command("ffmpeg", params...)
	if err != nil {
		log.E("TYPE_FLV exec command with input: %v, output: %v, params: %v error: %v", input, output, params, err)
		return go_trans.TransCommandError, go_trans.TransMessage{}, go_trans.Error{Err: err, Lines: []string{cmdOutput}}
	}
	end = util.Now13()
	log.D("TYPE_FLV exec command with input: %v, output: %v, params: %v success with cost: %v", input, output, params, end-beg)

	// return trans result
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
	log.D("TYPE_FLV command with input: %v, output: %v, params: %v success with message: %v", input, output, params, util.S2Json(message))

	return go_trans.TransOk, message, go_trans.Error{}
}

func (flv *Flv) Pid() int {
	if flv.Cmd == nil {
		return -1
	}
	return flv.Cmd.Process.Pid
}

func (flv *Flv) Cancel() error {
	if flv.Cmd == nil {
		return util.NewError("%v", go_trans.ErrorCode[go_trans.TransProcessNotExist])
	}

	return flv.Cmd.Process.Kill()
}

func (flv *Flv) Progress() (util.Map, error) {
	return nil, nil
}
