package go_trans

// Transcoding task
type Task struct {
	Id     string
	Input  string
	Output string
	TransCoder
}
