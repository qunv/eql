package mem

import (
	"github.com/qunv/eql/core/val"
	"sync"
)

type Mem interface {
	Set(key string, value val.EqlValue)
	Get(key string) val.EqlValue
	GetParent() Mem
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
		v := value.(val.EqlValue)
		return val.NewEqlValue(v.Value())
	}
	if m.parent != nil {
		return m.parent.Get(key)
	}
	return nil
}

func (m *_mem) GetParent() Mem {
	return m.parent
}
