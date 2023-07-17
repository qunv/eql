package action

import "github.com/qunv/eql/core/val"

type EqlInput [][]string

func (e EqlInput) Get(row int, col int) string {
	return e[row-1][col-1]
}

func (e EqlInput) Set(row int, col int, value string) {
	e[row-1][col-1] = value
}

type result struct {
	value val.EqlValue
	err   error
}
