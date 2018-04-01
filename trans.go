package go_trans

import (
	"github.com/tangs-drm/go-trans/util"
	"log"
	"path/filepath"
	"sync"
)

type TransPlugin interface {
	// Return the type of the transcode plug-in
	Type() string

	// Start the transcoding task.
	//
	// input: Input file name.
	// output: Output file name.
	// args: The parameters of the transcoding execution, such as
	//		{"-b:v": 1200000, "-r": 30}.
	//
	// string: The output information of the transcoding,
	// 		including the printing information of the transcoding success
	// 		and the failure of the transcoding.
	// error: Error information of the system.
	Exec(input, output string, args map[string]interface{}) (string, error)

	// Cancel the current transcoding task.
	//
	// error: error message.
	Cancel() error

	//
	Process() (map[string]interface{}, error)
}

// Transcoding task scheduler
type TransManage struct {
	MaxRunningNum  int
	CurrentRunning int

	Formats     []string
	TransPlugin map[string]TransPlugin
	Tasks       []*Task

	addSign  chan int
	execSign chan string
	isLoop   bool
	lock     *sync.Mutex
}

// The default number of transcoding threads
var DefaultMaxRunningNum = 1

// The default trans manager.
var DefaultTransManager = &TransManage{
	MaxRunningNum:  DefaultMaxRunningNum,
	CurrentRunning: 0,

	Formats:     []string{},
	TransPlugin: map[string]TransPlugin{},
	Tasks:       []*Task{},

	addSign:  make(chan int, 256),
	execSign: make(chan string, DefaultMaxRunningNum),
	isLoop:   false,
	lock:     &sync.Mutex{},
}

var DefaultFormats = []string{"flv"}

// Registering a supported transcode format with the transPlugin.
//
// format: video format like .flv, .avi.
// transPlugin: transcoding plugin.
//
// error: error message.
func (tm *TransManage) RegisterPlugin(format string, transPlugin TransPlugin) error {
	if _, ok := tm.TransPlugin[format]; ok {
		tm.TransPlugin[format] = transPlugin
	}
	tm.Formats = append(tm.Formats, format)
	tm.TransPlugin[format] = transPlugin
	return nil
}

func RegisterPlugin(format string, transPlugin TransPlugin) error {
	return DefaultTransManager.RegisterPlugin(format, transPlugin)
}

// GetFormats return the supported transcoding format
func (tm *TransManage) GetFormats() []string {
	return tm.Formats
}

func GetFormats() []string {
	return DefaultTransManager.GetFormats()
}

// SetMaxRunningNum set the maximum number of transcoding threads.This method
// is called if the call needs to be executed before method TransManage.Run().
func (tm *TransManage) SetMaxRunningNum(num int) {
	if num < 1 {
		num = 1
	}
	tm.execSign = make(chan string, num)
	tm.MaxRunningNum = num
}

func SetRuningNum(num int) {
	DefaultTransManager.SetMaxRunningNum(num)
}

// AddTask add a transcoding task, but just add the transcoding queue at this time,
// and do not really start transcoding.
//
// input: Input filename.
// output: Output filename.
func (tm *TransManage) AddTask(input, output string) (Task, error) {
	tm.lock.Lock()
	defer tm.lock.Unlock()

	// check input and output
	var inputExt = filepath.Ext(input)
	var outputExt = filepath.Ext(output)

	if "" == inputExt {
		return Task{}, util.Error("input is invalid: %v", input)
	}
	if "" == outputExt {
		return Task{}, util.Error("output is invalid: %v", output)
	}
	var plugin = tm.TransPlugin[inputExt]
	if plugin == nil {
		return Task{}, util.Error("unsupported format: %v", inputExt)
	}
	var task = &Task{
		Id:          util.UUID(),
		Input:       input,
		Output:      output,
		TransPlugin: plugin,
	}

	// todo. save into database.
	tm.Tasks = append(tm.Tasks, task)

	tm.addSign <- 1

	return *task, nil
}

func (tm *TransManage) StartTask() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("TransManage error: %v", err)
		}
	}()

}

// ListTask list the transcoding task.
//
// limit: Maximum number tasks return when func exec. If limit is less than 0, all of the task data is returned.
// skip: List start from skip.
//
// []Task: Tasks' detail.
// int: The count of all tasks.
func (tm *TransManage) ListTask(limit, skip int) ([]Task, int) {
	return nil, 0
}

func (tm *TransManage) Cancel(id string) error {
	return nil
}

func (tm *TransManage) Process(id []string) {

}

func (tm *TransManage) CallBack(id string) error {
	return nil
}
