package ioindexer

import (
	"fmt"
	"io/ioutil"
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
		}
	}
}

func readEmail(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	check(err)

	text := string(data)
	fmt.Println("texto: " + text)
	return text
}
