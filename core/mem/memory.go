package mem

import (
	"github.com/qunv/eql/core/val"
	"sync"
)

type Mem interface {
	Set(key string, value val.EqlValue)
	Get(key string) val.EqlValue
	GetParent() Mem
	Level() uint
}

type _mem struct {
	m      sync.Map
	parent Mem
	level  uint
}

func NewMemory(parent Mem) Mem {
	var level uint
	if parent == nil {
		level = 0
	} else {
		level = parent.Level() + 1
	}
	return &_mem{
		parent: parent,
		level:  level,
	}
}

func (m *_mem) Set(key string, value val.EqlValue) {
	level := m.Level()
	m.set(key, value, level)
}

func (m *_mem) set(key string, value val.EqlValue, level uint) {
	p := m.parent
	if p != nil {
		if p.Get(key) != nil {
			p.Set(key, value)
			return
		}
		p.(*_mem).set(key, value, level)
	}
	if m.Level() == level {
		m.m.Store(key, value)
	}
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

func (m *_mem) Level() uint {
	return m.level
}
