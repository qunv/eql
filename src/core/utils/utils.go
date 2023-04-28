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
