package test

import (
	"github.com/qunv/eql/core/val"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_Exec_LT(t *testing.T) {
	tests := []TestCase{
		{
			name: "Test LT number should return success",
			eql:  "LT(4, 2)",
			assert: func(value val.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "false", value.String())
			},
		},
		{
			name: "Test LT number and identify should return success",
			eql:  "LT(4, C1)",
			assert: func(value val.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "false", value.String())
			},
		},
		{
			name: "Test LT number and identify should return success",
			eql:  "LT(A1, C1)",
			assert: func(value val.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "true", value.String())
			},
		},
		{
			name: "Test LT number and identify should return success",
			eql:  "LT(D1, SUM(A1, A1:C1))",
			assert: func(value val.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "true", value.String())
			},
		},
		{
			name: "Test LT number and identify should return success",
			eql:  "LT(D1, SUM(A1,C1) + 1)",
			assert: func(value val.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "true", value.String())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := p.Exec(tt.eql)
			tt.assert(value, err)
		})
	}
}
