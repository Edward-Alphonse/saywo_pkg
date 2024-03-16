package storage

import (
	"encoding/json"
)

// Codec 使用go2cache需要实现下面的接口
type Codec interface {
	FormatKey(uint64) string //根据主键生成cache key
	Marshal(data any) ([]byte, error)
	Unmarshal(data []byte) (any, error)
}

type CacheWrap[T any] struct {
	Value T
}

type BaseCodec[T any] struct {
	Module    string
	Submodule []string
}

func (c *BaseCodec[T]) FormatKey(key uint64) string {
	generator := NewKeyGenerator(c.Module, c.Submodule, key)
	return generator.Generate()
}
func (c *BaseCodec[T]) Marshal(data any) ([]byte, error) {
	wrap := CacheWrap[any]{
		Value: data,
	}
	bs, err := json.Marshal(wrap)
	return bs, err
}

func (c *BaseCodec[T]) Unmarshal(data []byte) (any, error) {
	wrap := &CacheWrap[T]{}
	err := json.Unmarshal(data, wrap)
	return wrap.Value, err
}
