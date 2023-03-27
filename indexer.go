package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/miguelgz36/IndexerGolang/ioindexer"
)

func main() {

	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	nameFolderData := os.Args[1:][0]
	fmt.Println(nameFolderData)
	nameEmails := ioindexer.GetListOfEmails(nameFolderData)

	for _, nameEmail := range nameEmails {
		ioindexer.ReadEmails(nameFolderData, nameEmail)
	}

}
