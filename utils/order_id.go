package utils

import (
	"github.com/google/uuid"
	"strings"
)

// GenerateOrderNo 生成订单号
func GenerateOrderNo() string {
	u := uuid.New()
	return Md5(u.String())
}

func GenerateID32() (string, error) {
	u, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	id := strings.Replace(u.String(), "-", "", -1)
	return id, nil
}
