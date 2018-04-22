package network

import (
	"encoding/json"
	"github.com/tsdrm/go-trans"
	"github.com/tsdrm/go-trans/log"
	"github.com/tsdrm/go-trans/util"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
)

type RemoteTask struct {
	Input  string   `json:"input"`
	Path   string   `json:"path"`
	Format string   `json:"format"`
	Args   util.Map `json:"args"`
}

type Response struct {
	Code    int      `json:"code"`
	Error   string   `json:"error,omitempty"`
	Message string   `json:"message"`
	Data    util.Map `json:"data,omitempty"`
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	var err error
	var resp = Response{Data: util.Map{}}
	defer func() {
		_, err = w.Write([]byte(util.S2Json(resp)))
		if err != nil {
			log.E("[AddTask] Write data to w with resp: %v error: %v", util.S2Json(resp), err)
		}
	}()
	bys, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.E("[AddTask] Received bad request body: %v", err)
		resp.Code, resp.Error = go_trans.HTTPRequestBodyError, go_trans.ErrorCode[go_trans.HTTPRequestBodyError]
		resp.Message = err.Error()
		return
	}
	defer r.Body.Close()

	var rt = RemoteTask{}
	err = json.Unmarshal(bys, &rt)
	if err != nil {
		log.E("[AddTask] Trans body(%v) to RemoteTask error: %v", string(bys), err)
		resp.Code, resp.Error = go_trans.HTTPRequestBodyError, go_trans.ErrorCode[go_trans.HTTPRequestBodyError]
		resp.Message = err.Error()
		return
	}

	//
	if rt.Input == "" || rt.Format == "" {
		err = util.NewError("input/format is invalid")
		log.E("[AddTask] Received task: %v error: %v", util.S2Json(rt), err)
		resp.Code, resp.Error = go_trans.HTTPRequestParamsError, go_trans.ErrorCode[go_trans.HTTPRequestParamsError]
		resp.Message = err.Error()
		return
	}

	var newName = util.UUID() + rt.Format
	code, task, err := go_trans.AddTask(filepath.Join(rt.Path, rt.Input), newName, rt.Args)
	if err != nil {
		log.E("[AddTask] Add task with RemoteTask: %v error: %v", util.S2Json(rt), err)
		resp.Code, resp.Error = code, go_trans.ErrorCode[code]
		resp.Message = err.Error()
		return
	}

	resp.Code = go_trans.StatusOk
	resp.Message = go_trans.ErrorCode[go_trans.StatusOk]
	resp.Data["serverTime"] = util.Now13()
	resp.Data["task"] = task
	log.D("[AddTask] Add task with remoteTask: %v success with taskId: %v", util.S2Json(rt), task.Id)
}

func ListTasks(w http.ResponseWriter, r *http.Request) {
	var resp = Response{Data: util.Map{}}
	var err error

	defer func() {
		_, err = w.Write([]byte(util.S2Json(resp)))
		if err != nil {
			log.E("[ListTasks] Write data to w with resp: %v error: %v", util.S2Json(resp), err)
		}
	}()
	var page, pageCount int
	page, err = strconv.Atoi(r.FormValue("page"))
	if err != nil {
		log.E("[ListTasks] ListTasks received page error: %v", err)
		resp.Code, resp.Error = go_trans.HTTPRequestParamsError, go_trans.ErrorCode[go_trans.HTTPRequestParamsError]
		resp.Message = err.Error()
		return
	}
	pageCount, err = strconv.Atoi(r.FormValue("pageCount"))
	if err != nil {
		log.E("[ListTasks] ListTasks received pageCount error: %v", err)
		resp.Code, resp.Error = go_trans.HTTPRequestParamsError, go_trans.ErrorCode[go_trans.HTTPRequestParamsError]
		resp.Message = err.Error()
		return
	}

	var tasks, count = go_trans.ListTask(page, pageCount)

	resp.Code = go_trans.StatusOk
	resp.Message = go_trans.ErrorCode[go_trans.StatusOk]
	resp.Data["tasks"] = tasks
	resp.Data["count"] = count

	log.D("[ListTasks] ListTasks list with page: %v, pageCount: %v success with length: %v, count: %v", page, pageCount, len(tasks), count)
}

func Cancel(w http.ResponseWriter, r *http.Request) {
	var taskId = r.FormValue("taskId")
	var err error
	var code int
	var resp = Response{Data: util.Map{}}

	defer func() {
		_, err = w.Write([]byte(util.S2Json(resp)))
		if err != nil {
			log.E("[Cancel] Write data to w with resp: %v error: %v", util.S2Json(resp), err)
		}
	}()

	code, err = go_trans.Cancel(taskId)
	if err != nil {
		log.E("[Cancel] Cancel task with taskId: %v error: %v", taskId, err)
		resp.Code, resp.Error = code, go_trans.ErrorCode[code]
		resp.Message = err.Error()
		return
	}

	resp.Code = go_trans.StatusOk
	resp.Message = go_trans.ErrorCode[go_trans.StatusOk]
}
