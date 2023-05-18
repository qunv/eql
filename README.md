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

<details><summary>Math</summary>

<blockquote>
  <details><summary> SUM </summary>

Returns the sum of a series of numbers and/or cells

#### Sample usage

```shell
SUM(A2:A100)

SUM(1,2,3,4,5)

SUM(1,2,A2:A50)

SUM(1,SUM(A1:A3, 1), 2+A3)
```

#### Syntax

```shell
SUM(value1, [value2, ...])
```

- `value1` - The first number or range to add together.

- `value2`, ... - [ OPTIONAL ] - Additional numbers or ranges to add to value1.

#### Example

data

| A   | B   | C   |
|-----|-----|-----|
| 1   | 3   | 5   |
| 2   | 4   | 6   |


| Formula                | result |
|------------------------|--------|
| SUM(4, 2)              | 6      |
| SUM(A1, B1)            | 4      |
| SUM(A1:B2)             | 10     |
| SUM(A1:B2, SUM(3, C1)) | 18     |

  </details>

  <details><summary> ABS </summary>

Returns the absolute value of a number.

#### Sample usage

```shell
ABS(-2)

ABS(A2)
```

#### Syntax

```shell
ABS(value)
```

- `value` - The number of which to return the absolute value.

#### Example

| Formula | result   |
|---------|----------|
| ABS(-1) | 1        |
| ABS(A1) | 1        |

  </details>
</blockquote>
</details>

<details><summary>Operator</summary>

- ADD
- AVG
- CONCAT
- DIVIDE
- EQ
- MULTIPLY
- GT
- GTE
</details>

<details><summary> Logical </summary>
<blockquote>
  <details><summary> IF </summary>

Returns one value if a logical expression is `TRUE` and another if it is `FALSE`

#### Sample usage

```shell
IF(A2 = "foo","A2 is foo")

IF(A2,"A2 was true","A2 was false")

IF(TRUE,4,5)
```

#### Syntax

```shell
IF(logical_expression, value_if_true, value_if_false)
```

- `logical_expression` - An expression or reference to a cell containing an expression that represents some logical value, i.e. TRUE or FALSE.

- `value_if_true` - The value the function returns if logical_expression is TRUE.

- `value_if_false` - The value the function returns if logical_expression is FALSE.

#### Example

| Formula                  | result   |
|--------------------------|----------|
| IF(FALSE, 1, "false ne") | false ne |
| IF(TRUE=TRUE, 1, 2)      | 1        |
| IF(TRUE=TRUE, 1, 2)      | 1        |
| IF(1>SUM(1, 2), 1, 2)    | 2        |

  </details>
</blockquote>

</details>

# Contribute

This repo is required [antlr](https://www.antlr.org/) to define grammar, make sure install it, change your own `EqlLexer.g4`
and `EqlParser.g4` then run `make run` first

- Fork repository
- Create a feature branch
- Open a new pull request
- Create an issue for bug report or feature request

# License
The MIT License (MIT). Please see [LICENSE](LICENSE) for more information.