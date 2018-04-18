package go_trans

import (
	"github.com/tangs-drm/go-trans/log"
	"github.com/tangs-drm/go-trans/util"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func init() {
	RegisterPlugin(TYPE_MOCKPLUGIN, func() TransPlugin {
		return &MockPlugin{}
	})

	// create callback server
	httptest.NewRecorder()
	server := httptest.NewServer(nil)
	server.Start()
}

// create mock plugin
const TYPE_MOCKPLUGIN string = ".MOCKPLUGIN"

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
	time.Sleep(5 * time.Second)
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

func TestTransManage(t *testing.T) {
	// test formats
	formats := GetFormats()
	if len(formats) != 1 || formats[0] != TYPE_MOCKPLUGIN {
		t.Error(formats)
		return
	}
	// test set MaxRunningNum
	SetMaxRunningNum(3)
	if DefaultTransManager.MaxRunningNum != 3 {
		t.Error(DefaultTransManager.MaxRunningNum)
		return
	}

	var tasks []Task
	var task Task
	var num int
	var err error
	var args util.Map
	// list tasks
	tasks, num = ListTask(-1, 2)
	if num != 0 || len(tasks) != 0 {
		t.Error(num, tasks)
		return
	}

	// add task; invalid input
	task, err = AddTask("mockInput", "mockOutput.mp4", args)
	if err == nil || !strings.Contains(err.Error(), "input is invalid") {
		t.Error(err)
		return
	}
	task, err = AddTask("mockInput.flv", "mockOutput.mp4", args)
	if err == nil || !strings.Contains(err.Error(), "unsupported format") {
		t.Error(err)
		return
	}
	// add task; invalid output
	task, err = AddTask("mockInput."+TYPE_MOCKPLUGIN, "mockOutput", args)
	if err == nil || !strings.Contains(err.Error(), "output is invalid") {
		t.Error(err)
		return
	}

	// start task exec runner
	RunTask()

	// add task; normal
	task, err = AddTask("mockInput."+TYPE_MOCKPLUGIN, "mockOutput.mp4", args)
	if err != nil {
		t.Error(err)
		return
	}
	log.D("task -- %v", util.S2Json(task))
	if task.Status != TransNotStart {
		t.Error(util.S2Json(task))
		return
	}
	time.Sleep(time.Second)
}
