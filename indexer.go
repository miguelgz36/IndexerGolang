package main

import (
	"fmt"
	"os"

	"github.com/miguelgz36/IndexerGolang/ioindexer"
	//"github.com/miguelgz36/IndexerGolang/record"
)

func main() {
	nameFolderData := os.Args[1:][0]
	fmt.Println(nameFolderData)
	nameEmails := ioindexer.GetListOfEmails(nameFolderData)

	for i, nameEmail := range nameEmails {
		if i < 1 {
			ioindexer.ReadEmails(nameFolderData, nameEmail)
		}
	}

	//record.PostData(data)
}
