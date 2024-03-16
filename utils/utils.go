package utils

import "reflect"

func GetPtr[T comparable](p T) *T {
	tmp := p
	return &tmp
}

func IsNil(i any) bool {
	if i == nil {
		return true
	}
	defer func() {
		recover()
	}()
	vi := reflect.ValueOf(i)
	return vi.IsZero()
}
