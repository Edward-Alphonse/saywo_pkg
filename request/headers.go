package request

import (
	"encoding/json"
)

type Mapable interface {
	ToMap() map[string]any
}

type Parameter struct{}

func (p Parameter) ToMap() map[string]any {
	data := make(map[string]any)
	bytes, err := json.Marshal(p)
	if err != nil {
		return data
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return data
	}
	return data
}
