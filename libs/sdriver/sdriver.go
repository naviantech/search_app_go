package sdriver

import (
	"fmt"
)

type Instance struct {
	opts map[string]string
}

func (i *Instance) Traverse(key string) (val string) {
	return i[key]
}
