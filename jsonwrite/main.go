package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

const filePath = "./sample.json"
const noFilePath = "./no.json"
const filePath2 = "./sample2.json"
const noFilePath2 = "./no2.json"

const writeString = "AAAAAAAAAAAAAAA"

func main() {
	AppendJson()
	OverWrite()
}

func AppendJson() {
	isNew := false
	logFile, err := os.OpenFile(filePath, os.O_RDWR, 0600)
	if err != nil && err.Error() == fmt.Sprintf("open %s: no such file or directory", filePath) {
		isNew = true
		logFile, _ = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0600)
		_, err = logFile.Write([]byte("[]"))
	}
	defer logFile.Close()
	fi, err := logFile.Stat()
	leng := fi.Size()
	if isNew {
		_, err = logFile.WriteAt([]byte(fmt.Sprintf(`"%s"]`, writeString)), leng-1)
	} else {
		_, err = logFile.WriteAt([]byte(fmt.Sprintf(`,"%s"]`, writeString)), leng-1)
	}
}

func OverWrite() {
	logFile, err := os.OpenFile(filePath2, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		fmt.Print(err)
	}
	defer logFile.Close()

	buf := make([]byte, 1024*1024)
	for {
		n, err := logFile.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
	}
	// resize
	idx := bytes.IndexByte(buf, 0)
	buf = buf[:idx]

	var lst []string
	err = json.Unmarshal(buf, &lst)
	lst = append(lst, writeString)

	json_, err := json.Marshal(lst)
	logFile.WriteAt(json_, 0)
}
