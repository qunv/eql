package test

import (
	"github.com/qunv/eql/core/action"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_Exec_IF(t *testing.T) {
	tests := []TestCase{
		{
			name: "Test IF number should return success",
			eql:  "IF(FALSE, 1, \"false ne\")",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "false ne", value.String())
			},
		},
		{
			name: "Test IF number and identify should return success",
			eql:  "IF(FALSE, C1, B1)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "2", value.String())
			},
		},
		{
			name: "Test IF number and identify should return success",
			eql:  "IF(A1 > B1, C1, B1)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "2", value.String())
			},
		},
		{
			name: "Test IF number and identify should return success",
			eql:  "IF(D1 > SUM(A1, A1:C1), C1, 1)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "1", value.String())
			},
		},
		{
			name: "Test IF number and identify should return success",
			eql:  "IF(TRUE=TRUE, 1, 2)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "1", value.String())
			},
		},
		{
			name: "Test IF number and identify should return success",
			eql:  "IF(1=2, 1, 2)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "2", value.String())
			},
		},
		{
			name: "Test IF number and identify should return error",
			eql:  "IF(1+2, 1, 2)",
			assert: func(value action.EqlValue, err error) {
				assert.NotNil(t, err)
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
