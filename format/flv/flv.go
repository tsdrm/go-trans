package flv

type Flv struct {
	status string
}

const TYPE_FLV string = "flv"

func (flv *Flv) Type() string {
	return TYPE_FLV
}

func (flv *Flv) Exec(input, output string, args map[string]interface{}) (string, error) {
	return "", nil
}

func (flv *Flv) Cancel() error {
	return nil
}

func (flv *Flv) Process() (map[string]interface{}, error) {
	return nil, nil
}
