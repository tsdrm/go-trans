package go_trans

const (
	TransOk          = 0
	TransSystemError = 1
	TransTimeout     = 2
)

var ErrorCode = map[int]string{
	TransOk:          "Ok",
	TransSystemError: "System error",
	TransTimeout:     "Time out",
}
