package test

import (
	"github.com/qunv/eql/core/action"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_Exec_GTE(t *testing.T) {
	tests := []TestCase{
		{
			name: "Test GTE number should return success",
			eql:  "GTE(4; 2)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "true", value.String())
			},
		},
		{
			name: "Test GTE number and identify should return success",
			eql:  "GTE(4; C1)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "true", value.String())
			},
		},
		{
			name: "Test GTE number and identify should return success",
			eql:  "GTE(A1; C1)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "false", value.String())
			},
		},
		{
			name: "Test GTE number and identify should return success",
			eql:  "GTE(D1; SUM(A1; A1:C1))",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "false", value.String())
			},
		},
		{
			name: "Test GTE number and identify should return success",
			eql:  "GTE(D1; SUM(A1;C1) + 1)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "false", value.String())
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
