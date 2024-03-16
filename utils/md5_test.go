package utils

import (
	"fmt"
	"testing"
)

func TestPassword(t *testing.T) {
	password := Password("123456")
	fmt.Print(password)
}
