package network

import (
	"fmt"
	"github.com/tsdrm/go-trans"
	"github.com/tsdrm/go-trans/format/flv"
	"github.com/tsdrm/go-trans/log"
	"github.com/tsdrm/go-trans/util"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// define http router.
var router = map[string]http.HandlerFunc{
	"/addTask":   AddTask,
	"/listTasks": ListTasks,
	"/cancel":    Cancel,
}

func init() {
	go_trans.RegisterPlugin(".flv", func() go_trans.TransPlugin {
		return &flv.Flv{}
	})

	go_trans.RunTask()
}

func TestTask(t *testing.T) {

	var remoteTask RemoteTask
	var err error
	var result util.Map
	var taskId string
	var tasks []util.Map
	var count int
	// test for add task
	// normal
	remoteTask = RemoteTask{
		Input:  "../data/videos/f0.flv",
		Path:   "",
		Format: ".mp4",
		Args:   util.Map{},
	}
	result, err = HttpPost("/addTask", remoteTask)
	if err != nil {
		t.Error(err)
		return
	}
	log.D("result: %v", util.S2Json(result))
	taskId = result.Map("task").String("id")
	if "" == taskId {
		t.Error(util.S2Json(result))
		return
	}

	// list tasks
	result, err = HttpGet("/listTasks", "page=1&pageCount=10")
	if err != nil {
		t.Error(err)
		return
	}
	tasks = result.AryMap("tasks")
	count = result.Int("count")
	if count != 1 || len(tasks) != 1 {
		t.Error(err, util.S2Json(tasks))
		return
	}
	if tasks[0].String("id") != taskId {
		t.Error(util.S2Json(tasks))
		return
	}
	log.D("list task result: %v", util.S2Json(result))

	// add task unsupported format
	remoteTask = RemoteTask{
		Input:  "../data/videos/f0.abc",
		Path:   "",
		Format: ".mp4",
		Args:   util.Map{},
	}

	result, err = HttpPost("/addTask", remoteTask)
	if err == nil || !strings.Contains(err.Error(), "unsupported") {
		t.Error(err)
		return
	}

	time.Sleep(2 * time.Second)
	//time.Sleep(time.Hour)

	// cancel task
	result, err = HttpGet("/listTasks", "page=1&pageCount=10")
	if err != nil {
		t.Error(err)
		return
	}
	tasks = result.AryMap("tasks")
	count = result.Int("count")
	if count != 1 || len(tasks) != 1 {
		t.Error(err, util.S2Json(tasks))
		return
	}
	if tasks[0].String("id") != taskId {
		t.Error(util.S2Json(tasks))
		return
	}
	taskId = tasks[0].String("id")
	if taskId == "" {
		t.Error(util.S2Json(tasks[0]))
		return
	}
	result, err = HttpGet("/cancel", fmt.Sprintf("taskId=%v", taskId))
	if err != nil {
		t.Error(err)
		return
	}
	log.D("task cancel result: %v", util.S2Json(result))
	if result.Int("code") != 0 {
		t.Error(util.S2Json(result))
		return
	}

	// add task with empty body
	result, err = HttpPost("/addTask", util.Map{})
	if err == nil || !strings.Contains(err.Error(), "Request params error") {
		t.Error(err)
		return
	}

	// add task with invalid body
	{
		req, err := http.NewRequest("POST", "/addTask", strings.NewReader("abc"))
		if err != nil {
			t.Error(err)
			return
		}
		var recorder = httptest.NewRecorder()
		AddTask(recorder, req)
		err = util.Json2S(recorder.Body.String(), &result)
		if err != nil {
			t.Error(err)
			return
		}
		if result.Int("code") == 0 {
			t.Error(util.S2Json(result))
			return
		}
		if strings.Contains(result.String("message"), go_trans.ErrorCode[go_trans.HTTPRequestBodyError]) {
			t.Error(util.S2Json(result))
			return
		}
	}

	time.Sleep(10 * time.Second)

	// list tasks with invalid page
	result, err = HttpGet("/listTasks", "page=ab&pageCount=ab")
	if err != nil {
		t.Error(err)
		return
	}
	tasks = result.AryMap("tasks")
	count = result.Int("count")
	if len(tasks) != 0 || count != 0 {
		t.Error(count, len(tasks))
		return
	}

	// cancal task and not found
	result, err = HttpGet("/cancel", "taskId=akakak")
	if err == nil || !strings.Contains(err.Error(), go_trans.ErrorCode[go_trans.TransNotFound]) {
		t.Error(err)
		return
	}
}

func HttpGet(urlStr string, args string) (util.Map, error) {
	var req, err = http.NewRequest("GET", urlStr+"?"+args, nil)
	if err != nil {
		return nil, err
	}
	var recorder = httptest.NewRecorder()

	var handler = router[urlStr]
	if handler == nil {
		return nil, util.NewError("%v", "not found urlStr")
	}
	handler(recorder, req)

	var result util.Map
	err = util.Json2S(recorder.Body.String(), &result)
	if err != nil {
		return nil, err
	}
	if result.Int("code") != 0 {
		return nil, util.NewError("%v", util.S2Json(result))
	}
	return result.Map("data"), nil
}

func HttpPost(urlStr string, body interface{}) (util.Map, error) {
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(util.S2Json(body)))
	if err != nil {
		return nil, err
	}
	var recorder = httptest.NewRecorder()

	var handler = router[urlStr]
	if handler == nil {
		return nil, util.NewError("%v", "not found urlStr")
	}
	handler(recorder, req)

	var result util.Map
	err = util.Json2S(recorder.Body.String(), &result)
	if err != nil {
		return nil, err
	}
	if result.Int("code") != 0 {
		return nil, util.NewError("%v", util.S2Json(result))
	}
	return result.Map("data"), nil
}
