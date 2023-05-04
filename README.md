# eql

[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/qunv/eql)
![Build Status](https://github.com/qunv/eql/actions/workflows/test.yml/badge.svg?branch=main)

Sheet functions in [Go](http://www.golang.org) inspired by [Google sheet functions]()

[eql](https://github.com/qunv/eql) is designed to execute functions in sheet file in Go application.

# Install

```shell
go get github.com/qunv/eql
```

# Usage

```go
import (
	"github.com/qunv/eql"
)


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

func main() {
	input := initData()
	parser := eql.NewEqlParser(input)
	eql := "SUM(A1:B2; 1; AVG(A1:B2; C2); ADD(D3; 4))"
	result, err := parser.Exec(eql)
	if err != nil {
		// handle err
    }
	//hande result
}
```

# Functions

List of functions are supported currently, happy to contribute.

### Math

- SUM
- ABS

### Operator

- ADD
- AVG
- CONCAT
- DIVIDE
- EQ
- MULTIPLY

# Contribute

This repo is required [antlr](https://www.antlr.org/) to define grammar, make sure install it, change your own `EqlLexer.g4`
and `EqlParser.g4` then run `make run` first

- Fork repository
- Create a feature branch
- Open a new pull request
- Create an issue for bug report or feature request

# License
The MIT License (MIT). Please see [LICENSE](LICENSE) for more information.