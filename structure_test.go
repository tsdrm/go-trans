package go_trans

import (
	"fmt"
	"github.com/tangs-drm/go-trans/util"
	"strings"
	"testing"
)

func TestCall_ToString(t *testing.T) {
	var taskId = util.UUID()
	var call = Call{
		Code:         1,
		Error:        ErrorCode[TransCommandError],
		ErrorMessage: Error{Err: util.NewError("errorValue"), Lines: []string{"halo"}},
		Task: Task{
			Id:     taskId,
			Input:  "inputValue",
			Output: "outputValue",
			Status: "statusValue",
			Args:   util.Map{"argKey": 1},
			Plugin: &MockPlugin{},
		},
		Message: TransMessage{
			Input: InputFile{
				Cdn: "default",
			},
			Size:         100000,
			Position:     "positionValue",
			Cost:         12222,
			CreationTime: 99999,
			Duration:     1010,
		},
	}

	var callString = call.ToString()
	if !strings.Contains(callString, "errorValue") || !strings.Contains(callString, "halo") {
		t.Error(callString)
		return
	}
	if !strings.Contains(callString, taskId) || !strings.Contains(callString, "inputValue") || !strings.Contains(callString, "outputValue") {
		t.Error(callString)
		return
	}
	if !strings.Contains(callString, "statusValue") || !strings.Contains(callString, "argKey") {
		t.Error(callString)
		return
	}
	if !strings.Contains(callString, "default") || !strings.Contains(callString, "100000") || !strings.Contains(callString, "positionValue") {
		t.Error(callString)
		return
	}
	if !strings.Contains(callString, "12222") || !strings.Contains(callString, "99999") || !strings.Contains(callString, "1010") {
		t.Error(callString)
		return
	}
	fmt.Println("TestCall_ToString test success !!!")
}
