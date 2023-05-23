package main

import (
	"encoding/csv"
	"flag"
	"github.com/antlr4-go/antlr/v4"
	"github.com/qunv/eql/core"
	eqlantlr "github.com/qunv/eql/core/antlr"
	"log"
	"os"
	"regexp"
)

var eqlRegex = regexp.MustCompile("^[a-zA-z0-9].+\\.eql$")

func main() {

	subRun := flag.NewFlagSet("run", flag.ExitOnError)
	csvPath := subRun.String("csv", "", "data for our program")

	if len(os.Args) < 2 {
		log.Fatal("expected 'run' subcommands")
	}

	switch os.Args[1] {
	case "run":
		err := subRun.Parse(os.Args[3:])
		if err != nil {
			log.Fatal(err.Error())
		}
		executedFile := os.Args[2]
		if !eqlRegex.MatchString(executedFile) {
			log.Fatal("no eql file found")
		}
		if executedFile == "" {
			log.Fatal("no executed file")
		}
		if csvPath == nil || *csvPath == "" {
			log.Fatal("no csv file")
		}
		data := initData(*csvPath)
		run(executedFile, data)
	default:
		log.Fatal()
	}

}

func run(eqlFile string, data [][]string) {
	input, err := antlr.NewFileStream(eqlFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	lexer := eqlantlr.NewEqlLexer(input)
	eqlParser := eqlantlr.NewEqlParser(antlr.NewCommonTokenStream(lexer, 0))
	tree := eqlParser.Program()

	listener := core.NewEqlInterpreter(data)

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
}

func initData(csvPath string) [][]string {
	// open file
	f, err := os.Open(csvPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer f.Close()
	records, _ := csv.NewReader(f).ReadAll()
	return records
}
