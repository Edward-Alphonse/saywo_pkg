package id_generator

import (
	"github.com/sony/sonyflake"
)

var _snowflake *sonyflake.Sonyflake

func init() {
	_snowflake = sonyflake.NewSonyflake(sonyflake.Settings{})
}

func NextID() (uint64, error) {
	id, err := _snowflake.NextID()
	return id, err
}
