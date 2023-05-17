package action

type Action interface {
	Evaluate(input EqlInput) (EqlValue, error)
}
