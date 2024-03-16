package storage

import (
	"fmt"
	"strings"
)

type KeyGenerator struct {
	module     string
	subModules []string
	customs    []any
}

func NewKeyGenerator(module string, subModules []string, customs ...any) *KeyGenerator {
	if module == "" {
		panic("module is empty")
	}
	return &KeyGenerator{
		module:     module,
		subModules: subModules,
		customs:    customs,
	}
}

func (g *KeyGenerator) Generate() string {
	key := g.module
	if len(g.subModules) > 0 {
		key += ":" + strings.Join(g.subModules, ".")
	}
	if len(g.customs) > 0 {
		str := ""
		for i, v := range g.customs {
			if i < len(g.customs)-1 {
				str += fmt.Sprintf("%v.", v)
			} else {
				str += fmt.Sprintf("%v", v)
			}
		}
		if str != "" {
			key += ":" + str
		}
	}
	return key
}
