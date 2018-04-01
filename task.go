package go_trans

const TASK_RUNNING = "running"
const TASK_WAITING = "waiting"

// Transcoding task
type Task struct {
	Id     string
	Input  string
	Output string
	Status string
	TransPlugin
}
