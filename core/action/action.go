package action

type Action interface {
	Calculate(input EqlInput) (EqlValue, error)
}
