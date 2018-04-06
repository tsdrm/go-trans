package go_trans

const (
	TransOk              = 0
	TransSystemError     = 1
	TransTimeout         = 2
	TransCommandError    = 3
	TransNotFound        = 4
	TransTooManyTimes    = 5
	TransProcessNotExist = 6
)

var ErrorCode = map[int]string{
	TransOk:              "Ok",
	TransSystemError:     "System error",
	TransTimeout:         "Time out",
	TransCommandError:    "Command error",
	TransNotFound:        "Not found",
	TransTooManyTimes:    "Too many times",
	TransProcessNotExist: "Process not exist",
}
