package ioindexer

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/miguelgz36/IndexerGolang/record"
)

func check(er error) {
	if er != nil {
		panic(er)
	}
}

func checkFile(er error, file *os.File) {
	if er != nil {
		file.Close()
		panic(er)
	}
}

func GetListOfEmails(nameFolderData string) []string {
	emailsFileInfo, err := ioutil.ReadDir("./data/" + nameFolderData + "/maildir")
	check(err)

	nameEmailsFolders := make([]string, 0)
	for _, entry := range emailsFileInfo {
		if entry.IsDir() {
			nameEmailsFolders = append(nameEmailsFolders, entry.Name())
		}
	}
	return nameEmailsFolders
}

func ReadEmails(nameFolderData string, path string) {
	dir := "./data/" + nameFolderData + "/maildir/" + path
	nameEmailsSubFolders, err := ioutil.ReadDir(dir)
	check(err)

	for _, subDir := range nameEmailsSubFolders {
		if subDir.IsDir() {
			ReadEmails(nameFolderData, path+"/"+subDir.Name())
		} else {
			readEmail(dir + "/" + subDir.Name())
		}
	}
}

func convertFromMapToJson(mapToConvert map[string]string) {

	jsonData, err := json.Marshal(mapToConvert)
	check(err)
	record.PostData(string(jsonData))
}

func readEmail(filePath string) {

	file, err := os.Open(filePath)
	checkFile(err, file)

	scanner := bufio.NewScanner(file)

	mapOfProperties := map[string]string{}
	readingParams := true
	previousParameter := ""

	for scanner.Scan() {
		line := scanner.Text()

		if readingParams {
			indexFirstSeparator := strings.Index(line, ":")
			if indexFirstSeparator < len(line)-1 && indexFirstSeparator > 1 {
				key := line[:indexFirstSeparator]
				value := strings.Replace(line[indexFirstSeparator+1:], " ", "", 1)
				previousParameter = key
				mapOfProperties[key] = value
			} else {
				mapOfProperties[previousParameter] = mapOfProperties[previousParameter] + "\n" + line
			}
			if strings.Contains(line, "X-FileName") {
				readingParams = false
			}
		} else {
			mapOfProperties["message"] = mapOfProperties["message"] + "\n" + line
		}
	}

	check(err)
	file.Close()

	convertFromMapToJson(mapOfProperties)
}
