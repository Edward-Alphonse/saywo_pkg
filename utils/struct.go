package utils

import (
	"bytes"
	"encoding/json"
)

func Trans2Map(model any) (map[string]any, error) {
	data, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}
	dict := make(map[string]any)
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	err = d.Decode(&dict)
	if err != nil {
		return nil, err
	}
	return dict, nil
}

func Trans2Struct[T any](dict map[string]any) (*T, error) {
	data, err := json.Marshal(dict)
	if err != nil {
		return nil, err
	}
	model := new(T)
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	err = d.Decode(&model)
	if err != nil {
		return nil, err
	}
	return model, nil
}
