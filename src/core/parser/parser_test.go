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
		{1, 2, "test", "test"},
	}
	p := NewEqlParser(input)
	result := p.Exec("SUM(A0:B2; (C0 + 1)*2; AVG(C1; D0); CONCAT(A2; B2)) + SUM(A0; B0) + ADD(A0;B2)")
	f, err := result.Float64()
	assert.Equal(t, 38.5, f)

	//Add
	result = p.Exec("ADD(A0;B2)")
	f, err = result.Float64()
	assert.Nil(t, err)
	assert.Equal(t, float64(3), f)

	//Eq
	result = p.Exec("EQ(C3;D3)")
	assert.Equal(t, "true", result.String())

	result = p.Exec("CONCAT(C3;D3)")
	assert.Equal(t, "testtest", result.String())

	result = p.Exec("MULTIPLY(B0;C0)")
	assert.Equal(t, "6", result.String())

	result = p.Exec("DIVIDE(D0;B0)")
	assert.Equal(t, "2", result.String())
}
