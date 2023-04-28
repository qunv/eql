package utils

import (
	"github.com/qunv/eql/src/core/antlr"
	"strconv"
)

func GetColId(alias rune) int {
	return int(alias - 65)
}

func GetRowAndColum(ctx antlr.IMagicContext) (int, int) {
	col := []rune(ctx.IDENTIFIER().GetText())
	row, _ := strconv.Atoi(ctx.NUMBER().GetText())
	return row, GetColId(col[0])
}
