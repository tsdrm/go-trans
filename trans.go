package go_trans

import "sync"

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
}

// Transcoding task scheduler
type TransManage struct {
	lock *sync.Mutex
}

// AddTask add a transcoding task, but just add the transcoding queue at this time,
// and do not really start transcoding.
//
// input: Input filename.
// output: Output filename.
func (tm *TransManage) AddTask(intput, output string) {
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
