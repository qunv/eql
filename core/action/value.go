package action

import (
	"fmt"
	"github.com/qunv/eql/core/utils"
	"reflect"
)

type EqlValue interface {
	Value() interface{}
	Float64() (float64, error)
	String() string
	IsBool() bool
	Add(e EqlValue) error
	Minus(e EqlValue) error
	Mul(e EqlValue) error
	Div(e EqlValue) error
}

type value struct {
	v interface{}
	t reflect.Type
}

func NewEqlValue(v interface{}) EqlValue {
	return &value{v, reflect.TypeOf(v)}
}

func (v *value) Value() interface{} {
	return v.v
}

func (v *value) Float64() (float64, error) {
	return utils.GetFloat(v.v)
}

func (v *value) String() string {
	return fmt.Sprintf("%v", v.v)
}

func (v *value) Add(e EqlValue) error {
	val, err := e.Float64()
	if err != nil {
		return err
	}
	old, err := v.Float64()
	if err != nil {
		return err
	}
	v.v = old + val
	return nil
}

func (v *value) Minus(e EqlValue) error {
	val, err := e.Float64()
	if err != nil {
		return err
	}
	old, err := v.Float64()
	if err != nil {
		return err
	}
	v.v = old - val
	return nil
}

func (v *value) Mul(e EqlValue) error {
	val, err := e.Float64()
	if err != nil {
		return err
	}
	old, err := v.Float64()
	if err != nil {
		return err
	}
	v.v = old * val
	return nil
}

func (v *value) Div(e EqlValue) error {
	val, err := e.Float64()
	if err != nil {
		return err
	}
	old, err := v.Float64()
	if err != nil {
		return err
	}
	v.v = old / val
	return nil
}

func (v *value) IsBool() bool {
	return v.t.Kind() == reflect.Bool
}
