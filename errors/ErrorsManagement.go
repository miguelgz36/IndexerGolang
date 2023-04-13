package errors

import "os"

func Check(er error) {
	if er != nil {
		panic(er)
	}
}

func CheckFile(er error, file *os.File) {
	if er != nil {
		file.Close()
		panic(er)
	}
}
