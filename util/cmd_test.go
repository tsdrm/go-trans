package util

import (
	"github.com/tsdrm/go-trans/log"
	"testing"
)

func TestNewCmder(t *testing.T) {
	var cmder *Cmder
	// cmd normal
	cmder = NewCmder()
	output, err := cmder.Command("ping", "www.baidu.com")
	if err != nil {
		t.Error(err)
		return
	}
	log.D("cmder ping www.baidu.com \n: %v", output)

	// timeout
	cmder = NewCmder()
	cmder.SetTimeout(50)
	output, err = cmder.Command("ping", "www.baidu.com")
	if err == nil && err.Error() != "time out" {
		t.Error(err)
		return
	}
	log.D("cmder ping www.baidu.com time out")
}
