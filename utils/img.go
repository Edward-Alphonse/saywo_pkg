package utils

import (
	"fmt"
	"strings"
)

func RemoveBase64Header(base64Data string) string {
	// 查找 ";base64," 的位置
	base64Index := strings.Index(base64Data, ";base64,")
	if base64Index == -1 {
		fmt.Println("Base64 header not found in the data")
		return base64Data
	}

	// 截取头部信息之后的数据
	return base64Data[base64Index+len(";base64,"):]
}
