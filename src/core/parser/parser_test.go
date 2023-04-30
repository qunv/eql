package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_Handle(t *testing.T) {
	input := [][]interface{}{
		//A, B, C, D
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}
	p := NewEqlParser(input)
	result := p.Exec("SUM(A0:B2; (C0 + 1)*2; AVG(C1; D0); SUM(A2; B2)) + SUM(A0; B0) + ADD(A0;B2)")
	f, err := result.Float64()
	assert.Equal(t, 29.5, f)
	result = p.Exec("ADD(A0;B2)")
	f, err = result.Float64()
	assert.Nil(t, err)
	assert.Equal(t, float64(3), f)
}
