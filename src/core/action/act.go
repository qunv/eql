package action

type Action interface {
	Calculate(input [][]interface{}) (EqlValue, error)
}
