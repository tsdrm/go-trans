package network

import (
	"github.com/tsdrm/go-trans/log"
	"github.com/tsdrm/go-trans/util"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTask(t *testing.T) {

	var remoteTask RemoteTask
	// test for add task
	{
		remoteTask = RemoteTask{
			Input: "../data/videos/fo.flv",
			Path:  "",
			Args:  util.Map{},
		}
		var req, err = http.NewRequest("GET", "/addTask", strings.NewReader(util.S2Json(remoteTask)))
		if err != nil {
			t.Error(err)
			return
		}
		var recorder = httptest.NewRecorder()
		AddTask(recorder, req)
		if err != nil {
			t.Error(err)
			return
		}
		bys, err := ioutil.ReadAll(recorder.Body)
		if err != nil {
			t.Error(err)
			return
		}
		log.D("%v, %v", recorder.Code, string(bys))
	}
}
