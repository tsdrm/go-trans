package go_trans

import (
	"encoding/json"
	"fmt"
	"github.com/tsdrm/go-trans/log"
	"github.com/tsdrm/go-trans/util"
	"io/ioutil"
	"net/http"
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
}

// create mock plugin
const TYPE_MOCKPLUGIN string = ".MOCKPLUGIN"

var isMockError bool
var isCancel bool
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
		return TransCommandError, message, Error{Err: MockError}
	}
	log.D("MockPlugin start exec with isCancel: %v and will complete after 5 seconds.", isCancel)
	time.Sleep(5 * time.Second)
	log.D("MockPlugin start exec with isCancel: %v and will return", isCancel)
	if isCancel {
		isCancel = false
		return TransSystemError, message, Error{Err: util.NewError("already cancel")}
	}
	return StatusOk, message, Error{}
}

func (mp *MockPlugin) Cancel() error {
	log.D("MockPlugin cancel start with mock: %v", isMockError)
	isCancel = true
	if isMockError {
		return MockError
	}
	return nil
}

func (mp *MockPlugin) Progress() (util.Map, error) {
	return util.Map{}, nil
}

func (mp *MockPlugin) Pid() int {
	log.D("MockPlugin pid with mock: %v", isMockError)
	if isMockError {
		return -1
	}
	return 1
}

var callBackResult util.Map
var callBackCount int
var CallbackSuccess = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "success")
	callBackCount++
	bys, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.E("CallbackSuccess fail")
		callBackResult = util.Map{}
		return
	}
	defer r.Body.Close()
	err = json.Unmarshal(bys, &callBackResult)
	if err != nil {
		log.E("CallbackSuccess unmarshal data: %v error: %v", string(bys), err)
		callBackResult = util.Map{}
		return
	}
	log.D("CallbackSuccess success")
})
var CallbackFail = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	defer w.WriteHeader(http.StatusNotFound)
	callBackCount++
	bys, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.E("CallbackFail fail")
		callBackResult = util.Map{}
		return
	}
	defer r.Body.Close()
	err = json.Unmarshal(bys, &callBackResult)
	if err != nil {
		log.E("CallbackFail unmarshal data: %v error: %v", string(bys), err)
		callBackResult = util.Map{}
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
	var ts *httptest.Server

	var tasks []Task
	var task Task
	var count int
	var err error
	var args util.Map
	// list tasks
	tasks, count = ListTask(-1, -1)
	if count != 0 || len(tasks) != 0 {
		t.Error(count, tasks)
		return
	}

	// add task; invalid input
	_, task, err = AddTask("mockInput", "mockOutput.mp4", args)
	if err == nil || !strings.Contains(err.Error(), "input is invalid") {
		t.Error(err)
		return
	}
	_, task, err = AddTask("mockInput.flv", "mockOutput.mp4", args)
	if err == nil || !strings.Contains(err.Error(), "unsupported format") {
		t.Error(err)
		return
	}
	// add task; invalid output
	_, task, err = AddTask("mockInput."+TYPE_MOCKPLUGIN, "mockOutput", args)
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
	_, task, err = AddTask("mockInput."+TYPE_MOCKPLUGIN, "mockOutput.mp4", args)
	if err != nil {
		t.Error(err)
		return
	}
	log.D("task -- %v", util.S2Json(task))
	if task.Status != TransNotStart {
		t.Error(util.S2Json(task))
		return
	}

	time.Sleep(2 * time.Second)
	// list tasks empty
	tasks, count = ListTask(5, 15)
	if count != 1 || len(tasks) != 0 {
		t.Error(count, util.S2Json(tasks))
		return
	}

	// list tasks normal
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
		_, task, err = AddTask("mockInput."+TYPE_MOCKPLUGIN, "mockOutput.mp4", args)
		if err != nil {
			t.Error(err)
			return
		}
		_, err = Cancel(task.Id)
		if err != nil {
			t.Error(err)
			return
		}
		isCancel = false
		// check
		tasks, count = ListTask(1, 10)
		if count != 0 || len(tasks) != 0 {
			t.Error(count, util.S2Json(tasks))
			return
		}
	}

	{

		// test callback success
		ts = httptest.NewServer(CallbackSuccess)
		SetCallbackAddress(ts.URL)
		_, task, err = AddTask("mockInput."+TYPE_MOCKPLUGIN, "mockOutput.mp4", args)
		if err != nil {
			t.Error(err)
			return
		}
		time.Sleep(6 * time.Second)
		log.D("callbackResult: %v", util.S2Json(callBackResult))
		if callBackResult.Int("code") != 0 || callBackResult.Map("task").String("id") != task.Id {
			t.Error(util.S2Json(callBackResult), util.S2Json(task))
			return
		}

		// test callback success and trans error
		isMockError = true
		_, task, err = AddTask("mockInput."+TYPE_MOCKPLUGIN, "mockOutput.mp4", args)
		if err != nil {
			t.Error(err)
			return
		}
		time.Sleep(6 * time.Second)
		isMockError = false
		if callBackResult.Int("code") != TransCommandError || callBackResult.Map("task").String("id") != task.Id {
			t.Error(util.S2Json(callBackResult), util.S2Json(task))
			return
		}

		// test callback fail
		callBackCount = 0
		ts = httptest.NewServer(CallbackFail)
		SetCallbackAddress(ts.URL)
		_, task, err = AddTask("mockInput."+TYPE_MOCKPLUGIN, "mockOutput.mp4", args)
		if err != nil {
			t.Error(err)
			return
		}
		time.Sleep(30 * time.Second) // because of callback retry.
		if callBackCount != 3 {
			t.Error(callBackCount)
			return
		}
		if callBackResult.Int("code") != 0 || callBackResult.Map("task").String("id") != task.Id {
			t.Error(util.S2Json(callBackResult))
			return
		}
		time.Sleep(time.Second)

		// test callback with invalid callback address.
		SetCallbackAddress("invalidUrl")
		_, task, err = AddTask("mockInput"+TYPE_MOCKPLUGIN, "mockOutput.mp4", args)
		if err != nil {
			t.Error(err)
			return
		}
		time.Sleep(33 * time.Second)
		tasks, count = ListTask(1, -1)
		if count != 0 || len(tasks) != 0 {
			t.Error(count, util.S2Json(tasks))
			return
		}
	}

	// test cancel task error
	{
		ts = httptest.NewServer(CallbackSuccess)
		SetCallbackAddress(ts.URL)

		callBackCount = 0
		_, task, err = AddTask("mockInput2222"+TYPE_MOCKPLUGIN, "mockOutput.mp4", args)
		if err != nil {
			t.Error(err)
			return
		}
		time.Sleep(time.Second)
		isMockError = true
		_, err = Cancel(task.Id)
		if err == nil || err.Error() != MockError.Error() {
			t.Error(err)
			return
		}
		isMockError = false
		// check
		tasks, count = ListTask(1, 10)
		if count != 1 || len(tasks) != 1 {
			t.Error(count, util.S2Json(tasks))
			return
		}
		if tasks[0].Status != TransRunning {
			t.Error(util.S2Json(tasks[0]))
			return
		}
		task = tasks[0]

		time.Sleep(6 * time.Second)
		tasks, count = ListTask(1, 10)
		if count != 0 || len(tasks) != 0 {
			t.Error(count, util.S2Json(tasks))
			return
		}
		if callBackCount != 1 {
			t.Error(callBackCount)
			return
		}
		if callBackResult.Int("code") != TransSystemError || callBackResult.Map("task").String("id") != task.Id {
			t.Error(util.S2Json(callBackResult))
			return
		}

		// test cancel and not found
		_, err = Cancel("taskId")
		if err == nil || err.Error() != ErrorCode[TransNotFound] {
			t.Error(err)
			return
		}
	}

	// register plugin again
	RegisterPlugin(TYPE_MOCKPLUGIN, func() TransPlugin {
		return &MockPlugin{}
	})
	if len(DefaultTransManager.Formats) != 1 || len(DefaultTransManager.TransPlugin) != 1 {
		t.Error(DefaultTransManager.Formats)
		return
	}
	if DefaultTransManager.Formats[0] != TYPE_MOCKPLUGIN {
		t.Error(DefaultTransManager.Formats[0])
		return
	}
	if _, ok := DefaultTransManager.TransPlugin[TYPE_MOCKPLUGIN]; !ok {
		t.Error("key")
		return
	}

	// process
	Process(nil)
}
