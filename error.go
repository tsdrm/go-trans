package go_trans

const (
	// Transcode
	TransOk              = 0
	TransSystemError     = 1
	TransTimeout         = 2
	TransCommandError    = 3
	TransNotFound        = 4
	TransTooManyTimes    = 5
	TransProcessNotExist = 6

	// Http
	HTTPRequestBodyError   = 100
	HTTPRequestParamsError = 101
)

var ErrorCode = map[int]string{
	TransOk:                "Ok",
	TransSystemError:       "System error",
	TransTimeout:           "Time out",
	TransCommandError:      "Command error",
	TransNotFound:          "Not found",
	TransTooManyTimes:      "Too many times",
	TransProcessNotExist:   "Process not exist",
	HTTPRequestBodyError:   "Request body error",
	HTTPRequestParamsError: "Request params error",
}
