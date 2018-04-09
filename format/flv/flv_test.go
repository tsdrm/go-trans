package flv

import (
	"github.com/tangs-drm/go-trans"
	"github.com/tangs-drm/go-trans/log"
	"github.com/tangs-drm/go-trans/util"
	"testing"
)

func TestFlvTransCoding(t *testing.T) {
	var err error
	var f = &Flv{}

	if TYPE_FLV != f.Type() {
		t.Error(f.Type())
		return
	}

	if -1 != f.Pid() {
		t.Error(f.Pid())
		return
	}

	if err = f.Cancel(); err == nil || err.Error() != go_trans.ErrorCode[go_trans.TransProcessNotExist] {
		t.Error(err)
		return
	}

	var input, output string
	var args = util.Map{}

	// error with empty input
	code, message, err := f.Exec(input, output, args)
	if err == nil {
		t.Error(err)
		return
	}

	// success with empty args
	input = "../../data/videos/f0.flv"
	output = "../../data/output/" + util.UUID() + ".mp4"
	code, message, err = f.Exec(input, output, args)
	if err != nil {
		t.Error(err)
		return
	}
	log.D("code: %v, message: %v", code, util.S2Json(message))

	// success with resolution 1280*720
	args = util.Map{"-s": "1280*720"}
	input = "../../data/videos/f0.flv"
	output = "../../data/output/" + util.UUID() + ".mp4"
	code, message, err = f.Exec(input, output, args)
	if err != nil {
		t.Error(f.Lines)
		t.Error(err)
		return
	}
	log.D("code: %v, message: %v", code, util.S2Json(message))
}
