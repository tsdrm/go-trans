package go_trans

const TASK_RUNNING = "running"
const TASK_WAITING = "waiting"

// Transcoding task
type Task struct {
	Id     string                 `json:"id"`
	Input  string                 `json:"input"`
	Output string                 `json:"output"`
	Status string                 `json:"status"`
	Args   map[string]interface{} `json:"args"`
	Plugin TransPlugin            `json:"-"`
}

type Call struct {
	Code         int          `json:"code"`
	Error        string       `json:"error"`
	ErrorMessage error        `json:"errorMessage"`
	Task         Task         `json:"task"`
	Message      TransMessage `json:"message"`
}

type TransMessage struct {
	Input        InputFile `json:"input"`
	Size         int       `json:"size"`
	Position     string    `json:"position"`
	Cost         int       `json:"cost"`
	CreationTime int64     `json:"creationTime"`
	Duration     int       `json:"duration"`
}

type InputFile struct {
	Cdn string `json:"cdn"`
}
