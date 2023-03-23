package ioindexer

import (
	"bufio"
	"encoding/json"
	"fmt"
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
			fmt.Println("email en: " + dir)
			readEmail(dir + "/" + subDir.Name())
		}
	}
}

func convertFromMapToJson(mapToConvert map[string]string) {
	jsonBytes, err := json.MarshalIndent(mapToConvert, " ", "")
	check(err)

	jsonString := string(jsonBytes)

	fmt.Println("JSON:" + jsonString)
	record.PostData(jsonString)
}

func replaceRunes(r rune) rune {
	str := string(r)
	str = strings.Replace(str, "<", "(", -1)
	str = strings.Replace(str, ">", ")", -1)
	return []rune(str)[0]
}

func readEmail(filePath string) {

	file, err := os.Open(filePath)
	checkFile(err, file)

	scanner := bufio.NewScanner(file)

	mapOfProperties := map[string]string{}
	readingParams := true

	fmt.Println("READING:" + file.Name())

	for scanner.Scan() {
		line := strings.Map(replaceRunes, scanner.Text())
		fmt.Println("LINEA:" + line)

		if readingParams {
			indexFirstSeparator := strings.Index(line, ":")
			if indexFirstSeparator < len(line)-1 && indexFirstSeparator > 1 {
				key := line[:indexFirstSeparator]
				value := strings.Replace(line[indexFirstSeparator+1:], " ", "", 1)
				mapOfProperties[key] = value
			}
			if strings.Contains(line, "X-FileName") {
				readingParams = false
			}
		} else {
			mapOfProperties["message"] = mapOfProperties["message"] + line + "\n"
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("ERROR:" + err.Error())
	}
	file.Close()

	convertFromMapToJson(mapOfProperties)
}
