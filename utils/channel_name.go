package utils

import (
	"github.com/Edward-Alphonse/saywo_pkg/logs"
	"github.com/Edward-Alphonse/saywo_pkg/utils/id_generator"
	"strconv"
)

func GenChannelName() (string, error) {
	id, err := id_generator.NextID()
	if err != nil {
		logs.Error("gen id err", map[string]any{
			"err": err,
		})
		return "", err
	}
	channelName := strconv.Itoa(int(id))
	return channelName, nil
}
