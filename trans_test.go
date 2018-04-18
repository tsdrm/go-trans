package go_trans

import (
	"encoding/json"
	"fmt"
	"github.com/tangs-drm/go-trans/log"
	"github.com/tangs-drm/go-trans/util"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

func init() {
	RegisterPlugin(TYPE_MOCKPLUGIN, func() TransPlugin {
		return &MockPlugin{}
	})

	// create callback server
}

// create mock plugin
const TYPE_MOCKPLUGIN string = ".MOCKPLUGIN"

var isMockError bool
var MockError = util.NewError("%v", "Mock Error")

type MockPlugin struct {
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
	if isMockError {
		return TransCommandError, message, Error{Err: util.NewError("%s", "Mock Error")}
	}
	log.D("MockPlugin start exec and will complete after 5 seconds.")
	time.Sleep(5 * time.Second)
	return TransOk, message, Error{}
}

func (mp *MockPlugin) Cancel() error {
	if isMockError {
		return util.NewError("%s")
	}
	return nil
}

func (mp *MockPlugin) Progress() (util.Map, error) {
	return util.Map{}, nil
}

func (mp *MockPlugin) Pid() int {
	if isMockError {
		return -1
	}
	return 1
}

var callBackResult util.Map
var CallbackSuccess = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "success")
	bys, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.E("CallbackSuccess fail")
		callBackResult = nil
		return
	}
	defer r.Body.Close()
	err = json.Unmarshal(bys, &callBackResult)
	if err != nil {
		log.E("CallbackSuccess unmarshal data: %v error: %v", string(bys), err)
		callBackResult = nil
		return
	}
	log.D("CallbackSuccess success")
})
var CallbackFail = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	defer w.WriteHeader(http.StatusNotFound)
	bys, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.E("CallbackFail fail")
		callBackResult = nil
		return
	}
	defer r.Body.Close()
	err = json.Unmarshal(bys, &callBackResult)
	if err != nil {
		log.E("CallbackFail unmarshal data: %v error: %v", string(bys), err)
		callBackResult = nil
		return
	}
	log.D("CallbackFail not found")
})

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
	var count int
	var err error
	var args util.Map
	// list tasks
	tasks, count = ListTask(-1, 2)
	if count != 0 || len(tasks) != 0 {
		t.Error(count, tasks)
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

	// list tasks
	tasks, count = ListTask(-1, 2)
	if count != 0 || len(tasks) != 0 {
		t.Error(count, tasks)
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
	tasks, count = ListTask(-1, 2)
	if err != nil {
		t.Error(err)
		return
	}
	// check
	log.D("tasks %v, count: %v", util.S2Json(tasks), count)
	if count != 1 || len(tasks) != 1 {
		t.Error(util.S2Json(tasks), count)
		return
	}
	if tasks[0].Id != task.Id {
		t.Error(util.S2Json(tasks), util.S2Json(task))
		return
	}
	if tasks[0].Status != TransRunning {
		t.Error(util.S2Json(tasks[0]))
		return
	}

	time.Sleep(6 * time.Second)

	// list empty
	tasks, count = ListTask(1, 10)
	if count != 0 || len(tasks) != 0 {
		t.Error(util.S2Json(tasks), count)
		return
	}

	// test cancel task success
	{
		task, err = AddTask("mockInput."+TYPE_MOCKPLUGIN, "mockOutput.mp4", args)
		if err != nil {
			t.Error(err)
			return
		}
		err = Cancel(task.Id)
		if err != nil {
			t.Error(err)
			return
		}
		// check
		tasks, count = ListTask(1, 10)
		if count != 0 || len(tasks) != 0 {
			t.Error(count, util.S2Json(tasks))
			return
		}
	}
	// test cancel task error
	{
		task, err = AddTask("mockInput."+TYPE_MOCKPLUGIN, "mockOutput.mp4", args)
		if err != nil {
			t.Error(err)
			return
		}
		time.Sleep(time.Second)
		isMockError = true
		err = Cancel(task.Id)
		if err == nil || err.Error() != MockError.Error() {
			t.Error(err)
			return
		}
		isMockError = false
		// check
		tasks, count = ListTask(1, 10)
		if count != 0 || len(tasks) != 0 {
			t.Error(count, util.S2Json(tasks))
			return
		}
	}

	// setCallback
	//ts := httptest.NewServer(CallbackFail)
	//SetCallbackAddress(ts.URL)

}
