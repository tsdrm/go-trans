package flv

import "github.com/tangs-drm/go-trans"

type Flv struct {
	status string
}

const TYPE_FLV string = "flv"

func (flv *Flv) Type() string {
	return TYPE_FLV
}

func (flv *Flv) Exec(input, output string, args map[string]interface{}) (int, go_trans.TransMessage, error) {
	return go_trans.TransOk, go_trans.TransMessage{}, nil
}

func (flv *Flv) Cancel() error {
	return nil
}

func (flv *Flv) Progress() (map[string]interface{}, error) {
	return nil, nil
}
