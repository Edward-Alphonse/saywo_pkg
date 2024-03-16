package utils

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	MobileFormat1 = "%s-%s"
	MobileFormat2 = "+%s%s"
)

func FormatMobile(mobile, area, format string) string {
	if mobile == "" {
		return ""
	}
	v := strings.TrimPrefix(mobile, "0")
	if area == "" {
		area = "86"
	}

	if format == "" {
		format = MobileFormat1
	}
	return fmt.Sprintf(format, area, v)
}

func IsMobile(mb string) bool {
	b1, _ := regexp.MatchString(`^(12|13|14|15|16|17|18|19)\d{9}$`, mb)
	b2, _ := regexp.MatchString(`^(\d{1,4})-(\d{6,13})$`, mb)
	return b1 || b2
}

func VerifyMobile(mb, area string) (string, bool) {
	nmb := FormatMobile(mb, area, MobileFormat1)
	return nmb, IsMobile(nmb)
}
