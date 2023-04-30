package utils

import (
	"github.com/qunv/eql/src/core/antlr"
	"strconv"
)

func GetColId(alias rune) int {
	return int(alias - 65)
}

func GetRowAndColum(ctx antlr.IDefContext) (int, int) {
	col := []rune(ctx.IDENTIFIER().GetText())
	row, _ := strconv.Atoi(ctx.DIGIT().GetText())
	return row, GetColId(col[0])
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
