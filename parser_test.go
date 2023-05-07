package eql

import (
	"encoding/csv"
	"log"
	"os"
	"testing"

	"github.com/qunv/eql/core/action"

	"github.com/stretchr/testify/assert"
)

var input = initData()

var p = NewEqlParser(input)

type TestCase struct {
	name   string
	eql    string
	assert func(value action.EqlValue, err error)
}

func initData() [][]string {
	// open file
	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	records, _ := csv.NewReader(f).ReadAll()
	return records
}

func TestParser_Exec_ABS(t *testing.T) {
	tests := []TestCase{
		{
			name: "Test abs number should return success",
			eql:  "ABS(-1)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "1", value.String())
			},
		},
		{
			name: "Test abs with identify should return success",
			eql:  "ABS(A1)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "1", value.String())
			},
		},
		{
			name: "Test abs with wrong number should return error",
			eql:  "ABS(D4)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, value)
				assert.NotNil(t, err)
			},
		},
		{
			name: "Test abs with reach len params",
			eql:  "ABS(D4;A1)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, value)
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

func TestParser_Exec_ADD(t *testing.T) {
	tests := []TestCase{
		{
			name: "Test ADD number should return success",
			eql:  "ADD(-1; 2)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "1", value.String())
			},
		},
		{
			name: "Test ADD with identify should return success",
			eql:  "ADD(A1; B2)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "3", value.String())
			},
		},
		{
			name: "Test ADD with identify and a number",
			eql:  "ADD(A1; 1)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "2", value.String())
			},
		},
		{
			name: "Test ADD with reach limit len param",
			eql:  "ADD(A1; 1; 1)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, value)
				assert.NotNil(t, err)
			},
		},
		{
			name: "Test ADD wrong number",
			eql:  "ADD(A1; D4)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, value)
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

func TestParser_Exec_AVG(t *testing.T) {
	tests := []TestCase{
		{
			name: "Test AVG number should return success",
			eql:  "AVG(4; 2)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "3", value.String())
			},
		},
		{
			name: "Test AVG with identify should return success",
			eql:  "AVG(A1; B2)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "1.5", value.String())
			},
		},
		{
			name: "Test AVG with identify and a number",
			eql:  "AVG(A1; 1)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "1", value.String())
			},
		},
		{
			name: "Test ADD with multi param",
			eql:  "AVG(A1; 1; 4)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "2", value.String())
			},
		},
		{
			name: "Test AVG with input range",
			eql:  "AVG(A1:B2; A2)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "1.25", value.String())
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

func TestParser_Exec_CONCAT(t *testing.T) {
	tests := []TestCase{
		{
			name: "Test CONCAT number should return success",
			eql:  "CONCAT(4; 2)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "42", value.String())
			},
		},
		{
			name: "Test CONCAT number and identify should return success",
			eql:  "CONCAT(4; D4)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "4test", value.String())
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

func TestParser_Exec_DIVIDE(t *testing.T) {
	tests := []TestCase{
		{
			name: "Test DIVIDE number should return success",
			eql:  "DIVIDE(4; 2)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "2", value.String())
			},
		},
		{
			name: "Test DIVIDE number and identify should return success",
			eql:  "DIVIDE(D1; B1)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "2", value.String())
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

func TestParser_Exec_MULTIPLY(t *testing.T) {
	tests := []TestCase{
		{
			name: "Test MULTIPLY number should return success",
			eql:  "MULTIPLY(4; 2)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "8", value.String())
			},
		},
		{
			name: "Test MULTIPLY number and identify should return success",
			eql:  "MULTIPLY(D1; B1)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "8", value.String())
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

func TestParser_Exec_EQ(t *testing.T) {
	tests := []TestCase{
		{
			name: "Test EQ number should return success",
			eql:  "EQ(4; 2)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "false", value.String())
			},
		},
		{
			name: "Test EQ number and identify should return success",
			eql:  "EQ(C4; D4)",
			assert: func(value action.EqlValue, err error) {
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

func TestParser_Exec_SUM(t *testing.T) {
	tests := []TestCase{
		{
			name: "Test SUM number should return success",
			eql:  "SUM(4; 2)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "6", value.String())
			},
		},
		{
			name: "Test SUM number and identify should return success",
			eql:  "SUM(A1:B2; C1)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "9", value.String())
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

func TestParser_Exec_GT(t *testing.T) {
	tests := []TestCase{
		{
			name: "Test GT number should return success",
			eql:  "GT(4; 2)",
			assert: func(value action.EqlValue, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "true", value.String())
			},
		},
		{
			name: "Test GT number and identify should return success",
			eql:  "GT(C4; D4)",
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
