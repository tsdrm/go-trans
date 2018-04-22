package network

import (
	"encoding/json"
	"github.com/tsdrm/go-trans/log"
	"github.com/tsdrm/go-trans/util"
	"io/ioutil"
	"net/http"
)

type RemoteTask struct {
	Input string
	Path  string
	Args  util.Map
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	bys, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.E("[AddTask] received bad request body: %v", err)
		return
	}
	defer r.Body.Close()

	var rt = RemoteTask{}
	err = json.Unmarshal(bys, &rt)
	if err != nil {
		log.E("[AddTask] trans body(%v) to RemoteTask error: %v", string(bys), err)
		return
	}

	//
	if rt.Input == "" || rt.Path == "" {
	}
}
