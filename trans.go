package go_trans

import (
	"github.com/tangs-drm/go-trans/util"
	"path/filepath"
	"sync"
)

type TransCoder interface {
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
	Formats []TransCoder
	Format  map[string]TransCoder
	Tasks   []*Task

	lock *sync.Mutex
}

var DefaultFormats = []string{"flv"}

func (tm *TransManage) SetFormats(formats []string) error {
	return nil
}

func (tm *TransManage) RegisterFormats(format string, transCoder TransCoder) error {
	if _, ok := tm.Format[format]; ok {
		return util.Error("format: %v already exist", format)
	}
	tm.Formats = append(tm.Formats, transCoder)
	tm.Format[format] = transCoder
	return nil
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
	var transCoder = tm.Format[inputExt]
	if transCoder == nil {
		return Task{}, util.Error("unsupported format: %v", inputExt)
	}
	var task = &Task{
		Id:         util.UUID(),
		Input:      input,
		Output:     output,
		TransCoder: transCoder,
	}

	return *task, nil
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
