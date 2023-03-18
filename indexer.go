package main

import (
	"fmt"
	"os"

	"github.com/miguelgz36/IndexerGolang/ioindexer"
	"github.com/miguelgz36/IndexerGolang/record"
)

func main() {
	argsWithoutProg := os.Args[1:][0]
	fmt.Println(argsWithoutProg)
	data := ioindexer.GetDataBody()
	record.PostData(data)
}
