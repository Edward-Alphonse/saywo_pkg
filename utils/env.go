package utils

import (
	"os"
	"strings"
)

const (
	DEBUG_ENV   = "debug"   //在本地机器上开发
	RELEASE_ENV = "release" //线上环境
	TEST_ENV    = "test"    //在ecs环境下开发或者测试环境
)

var env = DEBUG_ENV

func init() {
	env = os.Getenv("run_mode")
	env = strings.ToLower(env)
	switch env {
	case RELEASE_ENV:
	case TEST_ENV:
	default:
		env = DEBUG_ENV
	}
}

func GetEnv() string {
	return env
}

func IsDevEnv() bool {
	return env == DEBUG_ENV || env == TEST_ENV
}
