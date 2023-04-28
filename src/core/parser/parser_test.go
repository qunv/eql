package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_Handle(t *testing.T) {
	input := [][]float64{
		//A, B, C, D
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}
	p := NewEqlParser(input)
	result := p.Exec("=SUM(A0:B2; C0; =AVG(C1; D0); =SUM(A2; B2)) + =SUM(A0; B0) + D3")
	assert.Equal(t, result, 25.5)
}
