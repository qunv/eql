package action

type EqlInput [][]interface{}

func (e EqlInput) Get(row int, col int) interface{} {
	return e[row-1][col-1]
}
