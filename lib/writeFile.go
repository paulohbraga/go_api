package lib

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteFile(text string, filename string, ip string) {

	execPath, err2 := os.Executable()
	check(err2)
	execFullPath := filepath.Dir(execPath)

	outputFilePath := filepath.Join(execFullPath, "/files/", filename)
	f, err := os.OpenFile(outputFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)

	defer func(f *os.File) {
		err := f.Close()
		check(err)
	}(f)

	err3 := f.Sync()
	check(err3)

	w := bufio.NewWriter(f)
	// create date and time string formated

	now := time.Now()
	formattedTime := now.Format("2006-01-02 15:04:05")

	linetoWrite := fmt.Sprint("IP Address: ", ip, " - Date: ", formattedTime, " message: ", text, "\n")
	_, err = w.WriteString(linetoWrite)
	check(err)

	errFlush := w.Flush()
	check(errFlush)

}
