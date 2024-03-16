package utils

import "encoding/json"

func Trans2Map(model any) (map[string]any, error) {
	bytes, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}
	dic := make(map[string]any)
	err = json.Unmarshal(bytes, &dic)
	if err != nil {
		return nil, err
	}
	return dic, nil
}

func Trans2Struct[T any](dict map[string]any) (*T, error) {
	bytes, err := json.Marshal(dict)
	if err != nil {
		return nil, err
	}
	model := new(T)
	err = json.Unmarshal(bytes, &model)
	if err != nil {
		return nil, err
	}
	return model, nil
}
