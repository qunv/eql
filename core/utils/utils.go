package utils

import (
	"errors"
	"github.com/qunv/eql/core/antlr"
	"strconv"
)

func GetColId(alias rune) int {
	return int(alias - 64)
}

func GetRowAndColum(ctx antlr.IDefContext) (int, int, error) {
	col := []rune(ctx.IDENTIFIER().GetText())
	row, _ := strconv.Atoi(ctx.DIGIT().GetText())
	if row <= 0 {
		return 0, 0, errors.New("row index not accept 0 or negative number")
	}
	return row, GetColId(col[0]), nil
}

func GetNumber(ctx antlr.INumberContext) float64 {
	f, _ := strconv.ParseFloat(ctx.DIGIT().GetText(), 64)
	if ctx.MINUS() != nil {
		return -f
	}
	return f
}

func GetFloat(unk interface{}) (float64, error) {
	switch i := unk.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint:
		return float64(i), nil
	case string:
		return strconv.ParseFloat(i, 64)
	default:
		panic("Type cannot to float64")
	}
}
