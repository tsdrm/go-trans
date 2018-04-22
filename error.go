package go_trans

const (
	StatusOk = 0

	// Transcode
	TransSystemError     = 1
	TransTimeout         = 2
	TransCommandError    = 3
	TransNotFound        = 4
	TransTooManyTimes    = 5
	TransProcessNotExist = 6
	TransParamsError     = 7

	// Http
	HTTPRequestBodyError   = 100
	HTTPRequestParamsError = 101
)

var ErrorCode = map[int]string{
	StatusOk: "Ok",

	// Transcode
	TransSystemError:     "System error",
	TransTimeout:         "Time out",
	TransCommandError:    "Command error",
	TransNotFound:        "Not found",
	TransTooManyTimes:    "Too many times",
	TransProcessNotExist: "Process not exist",
	TransParamsError:     "Params invalid",

	// Http
	HTTPRequestBodyError:   "Request body error",
	HTTPRequestParamsError: "Request params error",
}
