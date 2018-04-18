package go_trans

import "github.com/tangs-drm/go-trans/util"

const TASK_RUNNING = "running"
const TASK_WAITING = "waiting"

// Transcoding task
type Task struct {
	Id     string      `json:"id"`
	Input  string      `json:"input"`
	Output string      `json:"output"`
	Status string      `json:"status"`
	Args   util.Map    `json:"args"`
	Plugin TransPlugin `json:"-"`
}

type Call struct {
	Code         int          `json:"code"`
	Error        string       `json:"error"`
	ErrorMessage Error        `json:"errorMessage"`
	Task         Task         `json:"task"`
	Message      TransMessage `json:"message"`
}

func (call Call) ToString() string {
	var errorMessage = util.Map{}
	if call.ErrorMessage.Err != nil {
		errorMessage = util.Map{
			"err":   call.ErrorMessage.Err.Error(),
			"lines": call.ErrorMessage.Lines,
		}
	}
	var errorMessageMap util.Map
	err := util.Json2S(util.S2Json(call), &errorMessageMap)
	if err != nil {
		return ""
	}
	errorMessageMap["errorMessage"] = errorMessage
	return util.S2Json(errorMessageMap)
}

type TransMessage struct {
	Input        InputFile `json:"input"`        // input info
	Size         int       `json:"size"`         // input file size
	Position     string    `json:"position"`     // transcoding server name
	Cost         int       `json:"cost"`         // the time cost of transcoding
	CreationTime int64     `json:"creationTime"` // create time
	Duration     float32   `json:"duration"`     // the duration of input file
}

type InputFile struct {
	Cdn string `json:"cdn"`
}

// Error of transcoding.
type Error struct {
	Err   error    `json:"err"`   // error message returned from process. eg. exit status 1
	Lines []string `json:"lines"` // output message from stderr.
}
