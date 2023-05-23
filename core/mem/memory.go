package mem

import (
	"github.com/qunv/eql/core/val"
	"sync"
)

var store = &_mem{}

type _mem struct {
	m sync.Map
}

func Set(key string, value val.EqlValue) {
	store.m.Store(key, value)
}

func Get(key string) val.EqlValue {
	value, ok := store.m.Load(key)
	if ok {
		return value.(val.EqlValue)
	}
	return nil
}
