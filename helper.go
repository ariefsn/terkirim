package terkirim

import (
	"bytes"
	"encoding/json"
)

func convert[T any](from interface{}) (T, error) {
	var to T

	data, err := json.Marshal(from)

	if err != nil {
		return to, err
	}

	err = json.Unmarshal(data, &to)

	if err != nil {
		return to, err
	}

	return to, nil
}

func toJsonBody(v interface{}) (*bytes.Buffer, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(b), nil
}
