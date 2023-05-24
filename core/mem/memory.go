package mem

import (
	"github.com/qunv/eql/core/val"
	"sync"
)

type Mem interface {
	Set(key string, value val.EqlValue)
	Get(key string) val.EqlValue
	Parent() Mem
}

type _mem struct {
	m      sync.Map
	parent Mem
}

func NewMemory(parent Mem) Mem {
	return &_mem{
		parent: parent,
	}
}

func (m *_mem) Set(key string, value val.EqlValue) {
	m.m.Store(key, value)
}

func (m *_mem) Get(key string) val.EqlValue {
	value, ok := m.m.Load(key)
	if ok {
		return value.(val.EqlValue)
	}
	return nil
}

func (m *_mem) Parent() Mem {
	return m.parent
}
