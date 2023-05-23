package action

import "github.com/qunv/eql/core/val"

type Action interface {
	Evaluate(input EqlInput) (val.EqlValue, error)
}
