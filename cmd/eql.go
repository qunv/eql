package main

import (
	"encoding/csv"
	"github.com/qunv/eql"
	"log"
	"os"
	"regexp"
)

var eqlRegex = regexp.MustCompile("^[a-zA-z0-9].+\\.eql$")

func main() {

	//subRun := flag.NewFlagSet("run", flag.ExitOnError)
	//csvPath := subRun.String("csv", "", "data for our program")
	//
	//if len(os.Args) < 2 {
	//	log.Fatal("expected 'run' subcommands")
	//}

	//switch os.Args[1] {
	//case "run":
	//	err := subRun.Parse(os.Args[3:])
	//	if err != nil {
	//		log.Fatal(err.Error())
	//	}
	//	executedFile := os.Args[2]
	//	if !eqlRegex.MatchString(executedFile) {
	//		log.Fatal("no eql file found")
	//	}
	//	if executedFile == "" {
	//		log.Fatal("no executed file")
	//	}
	//	if csvPath == nil || *csvPath == "" {
	//		log.Fatal("no csv file")
	//	}
		data := initData("/home/joseph/Documents/self/eql/cmd/test.csv")
		parser := eql.NewEqlParser(data)
		if err := parser.Execf("/home/joseph/Documents/self/eql/cmd/main.eql"); err != nil {
			log.Fatal(err.Error())
		}
	//default:
	//	log.Fatal()
	//}

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
