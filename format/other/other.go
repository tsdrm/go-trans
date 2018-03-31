package other

type Other struct {
}

const TYPE_OTHER string = "other"

func (other *Other) Type() string {
	return TYPE_OTHER
}

func (other *Other) Exec(input, output string, args map[string]interface{}) (string, error) {
	return "", nil
}

func (other *Other) Cancel() error {
	return nil
}

func (other *Other) Process() (map[string]interface{}, error) {
	return nil, nil
}
