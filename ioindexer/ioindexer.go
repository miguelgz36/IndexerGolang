package ioindexer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func check(er error) {
	if er != nil {
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

func ReadEmails(nameFolderData string, path string, textEmails *[]string) {
	dir := "./data/" + nameFolderData + "/maildir/" + path
	nameEmailsSubFolders, err := ioutil.ReadDir(dir)
	check(err)

	for _, subDir := range nameEmailsSubFolders {
		if subDir.IsDir() {
			ReadEmails(nameFolderData, path+"/"+subDir.Name(), textEmails)
		} else {
			fmt.Println("email en: " + dir)
			text := readEmail(dir + "/" + subDir.Name())
			*textEmails = append(*textEmails, text)
			readText(text)
		}
	}
}

func customSplitFunc(r rune) bool {
	return r == '\n' || r == '\r'
}

func readText(text string) {
	lines := strings.FieldsFunc(text, customSplitFunc)
	indexMessage := 0
	mapOfProperties := map[string]string{}

	for indexLine, line := range lines {
		fmt.Println("LINEA:" + line)
		indexFirstSeparator := strings.Index(line, ":")
		if indexFirstSeparator < len(line)-1 && indexFirstSeparator > 1 {
			key := line[:indexFirstSeparator]
			value := strings.ReplaceAll(line[indexFirstSeparator+1:], " ", "")
			mapOfProperties[key] = value

			fmt.Printf("KEY: %s VALUE: %s\n", key, value)
		}
		isLastParam := strings.Contains(line, "X-FileName")
		if isLastParam {
			indexMessage = indexLine + 1
			break
		}
	}

	mapOfProperties["message"] = strings.Join(lines[indexMessage:], "")

	convertFromMapToJson(mapOfProperties)
}

func convertFromMapToJson(mapToConvert map[string]string) {
	jsonBytes, err := json.MarshalIndent(mapToConvert, " ", "")
	check(err)

	jsonString := string(jsonBytes)

	fmt.Println("JSON:" + jsonString)
}

func replaceRunes(r rune) rune {
	str := string(r)
	str = strings.Replace(str, "<", "(", -1)
	str = strings.Replace(str, ">", ")", -1)
	return []rune(str)[0]
}

func readEmail(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	check(err)

	text := string(data)
	text = strings.Map(replaceRunes, text)

	fmt.Println("TEXTO----: " + text)
	return text
}
