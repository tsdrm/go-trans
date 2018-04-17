package go_trans

import (
	"github.com/tangs-drm/go-trans/util"
	"testing"
)

const TYPE_MOCKPLUGIN string = "MOCKPLUGIN"

type MockPlugin struct {
	MockError bool
}

func (mp *MockPlugin) Type() string {
	return TYPE_MOCKPLUGIN
}

func (mp *MockPlugin) Exec(input, output string, args util.Map) (int, TransMessage, Error) {
	// return trans result
	var message = TransMessage{
		Input: InputFile{
			Cdn: "default",
		},
		Size:         2014,
		Position:     "",
		Cost:         20,
		CreationTime: util.Now13(),
		Duration:     10.32,
	}
	if mp.MockError {
		return TransCommandError, message, Error{Err: util.NewError("%s", "Mock Error")}
	}
	return TransOk, message, Error{}
}

func (mp *MockPlugin) Cancel() error {
	if mp.MockError {
		return util.NewError("%s", "Mock Error")
	}
	return nil
}

func (mp *MockPlugin) Progress() (util.Map, error) {
	return util.Map{}, nil
}

func (mp *MockPlugin) Pid() int {
	if mp.MockError {
		return -1
	}
	return 1
}

func init() {
	RegisterPlugin("mock", func() TransPlugin {
		return &MockPlugin{}
	})

}

func TestTransManage(t *testing.T) {

}
