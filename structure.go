package go_trans

const TASK_RUNNING = "running"
const TASK_WAITING = "waiting"

// Transcoding task
type Task struct {
	Id     string            `json:"id"`
	Input  string            `json:"input"`
	Output string            `json:"output"`
	Status string            `json:"status"`
	Args   map[string]string `json:"args"`
	Plugin TransPlugin       `json:"-"`
}

type Call struct {
	Code         int          `json:"code"`
	Error        string       `json:"error"`
	ErrorMessage error        `json:"errorMessage"`
	Task         Task         `json:"task"`
	Message      TransMessage `json:"message"`
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
	Err   error    // error message returned from process. eg. exit status 1
	Lines []string // output message from stderr.
}
