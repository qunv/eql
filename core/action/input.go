package action

type EqlInput [][]string

func (e EqlInput) Get(row int, col int) string {
	return e[row-1][col-1]
}
