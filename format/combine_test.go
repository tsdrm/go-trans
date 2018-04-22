package format

import (
	"fmt"
	"github.com/tsdrm/go-trans"
	"github.com/tsdrm/go-trans/format/flv"
	"testing"
)

func TestTransPlugin(t *testing.T) {
	Init()
	var plugin = go_trans.DefaultTransManager.TransPlugin["flv"]
	if plugin == nil {
		t.Error("plugin is nil")
		return
	}
	if flv.TYPE_FLV != plugin().Type() {
		t.Error(plugin().Type(), flv.TYPE_FLV)
		return
	}
	fmt.Println("TestTransPlugin test success")
}
