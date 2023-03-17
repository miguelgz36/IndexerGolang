package main

import (
	"github.com/miguelgz36/IndexerGolang/ioindexer"
	"github.com/miguelgz36/IndexerGolang/record"
)

func main() {
	data := ioindexer.GetDataBody()
	record.PostData(data)
}
