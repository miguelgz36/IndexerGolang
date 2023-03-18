package main

import (
	"fmt"
	"os"

	"github.com/miguelgz36/IndexerGolang/ioindexer"
	"github.com/miguelgz36/IndexerGolang/record"
)

func main() {
	nameFolderData := os.Args[1:][0]
	fmt.Println(nameFolderData)
	data := ioindexer.GetDataBody(nameFolderData)
	record.PostData(data)
}
