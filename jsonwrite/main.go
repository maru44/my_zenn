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

type Sample struct {
	Name    string `json:"name"`
	Old     int    `json:"old,omitempty"`
	Explain string `json:"explain"`
}

func main() {
}

func AppendJson(fileName string, object interface{}) {
	isNew := false
	file, err := os.OpenFile(fileName, os.O_RDWR, 0600)
	if err != nil && err.Error() == fmt.Sprintf("open %s: no such file or directory", fileName) {
		isNew = true
		file, _ = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0600)
		_, err = file.Write([]byte("[]"))
	}
	defer file.Close()
	fi, err := file.Stat()
	leng := fi.Size()

	json_, err := json.Marshal(object)

	if isNew {
		_, err = file.WriteAt([]byte(fmt.Sprintf(`%s]`, json_)), leng-1)
	} else {
		_, err = file.WriteAt([]byte(fmt.Sprintf(`,%s]`, json_)), leng-1)
	}
}

func OverWrite(fileName string, object interface{}) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	buf := make([]byte, 1024*1024)
	for {
		n, err := file.Read(buf)
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

	var lst []interface{}
	json.Unmarshal(buf, &lst)
	lst = append(lst, object)
	json_, _ := json.Marshal(lst)
	file.WriteAt(json_, 0)
}
