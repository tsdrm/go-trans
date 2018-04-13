package go_trans

import (
	"github.com/tangs-drm/go-trans/util"
	"testing"
)

const TYPE_MOCKPLUGIN string = "MOCKPLUGIN"

type MockPlugin struct {
}

func (mp *MockPlugin) Type() string {
	return TYPE_MOCKPLUGIN
}

func (mp *MockPlugin) Exec(input, output string, args util.Map) (int, TransMessage, Error) {
	return 0, TransMessage{}, Error{}
}

func (mp *MockPlugin) Cancel() error {
	return nil
}

func (mp *MockPlugin) Progress() (util.Map, error) {
	return util.Map{}, nil
}

func (mp *MockPlugin) Pid() int {
	return 1
}

func TestTransManage(t *testing.T) {

}
