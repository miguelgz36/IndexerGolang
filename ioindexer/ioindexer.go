package ioindexer

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/miguelgz36/IndexerGolang/errors"
	"github.com/miguelgz36/IndexerGolang/record"
)

var keysOfEmailToIndex = [4]string{"to", "from", "message", "subject"}

func GetListOfEmails(nameFolderData string) []string {
	emailsFileInfo, err := ioutil.ReadDir("./data/" + nameFolderData + "/maildir")
	errors.Check(err)

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
	errors.Check(err)

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
	errors.Check(err)
	record.PostData(string(jsonData))
}

func readEmail(filePath string) {
	time.Sleep(1 * time.Millisecond)

	go func() {

		file, err := os.Open(filePath)
		errors.CheckFile(err, file)

		scanner := bufio.NewScanner(file)

		mapOfProperties := map[string]string{}
		readingParams := true
		previousParameter := ""

		for scanner.Scan() {
			line := scanner.Text()
			if readingParams {

				indexFirstSeparator := strings.Index(line, ":")
				if indexFirstSeparator < len(line)-1 && indexFirstSeparator > 1 {
					key := strings.ToLower(line[:indexFirstSeparator])
					if containsKey(key) {
						value := strings.Replace(line[indexFirstSeparator+1:], " ", "", 1)
						previousParameter = key
						mapOfProperties[key] = value
					}
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
		file.Close()
		errors.Check(err)
		convertFromMapToJson(mapOfProperties)
	}()

}

func containsKey(searchString string) bool {
	for _, key := range keysOfEmailToIndex {
		if key == searchString {
			return true
		}
	}
	return false
}
