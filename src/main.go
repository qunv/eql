package main

import (
	"github.com/gin-gonic/gin"
	parser2 "github.com/qunv/eql/src/core/parser"
	"net/http"
)

type Request struct {
	Q    string      `json:"q"`
	Data [][]float64 `json:"data"`
}

type Response struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

func main() {
	router := gin.Default()
	router.POST("/exec", func(ctx *gin.Context) {
		var input Request
		err := ctx.ShouldBindJSON(&input)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, &Response{
				Error: err.Error(),
			})
		}
		parser := parser2.NewEqlParser(input.Data)
		result := parser.Exec(input.Q)
		ctx.JSON(http.StatusOK, &Response{
			Data: result,
		})
	})
	_ = router.Run("localhost:8080")
}
