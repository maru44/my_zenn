package jsonwrite_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// run
// go test -bench . -benchmem

// 要素を追加 (既存)
func BenchmarkAppend_AppendJson(b *testing.B) {
	b.ResetTimer()
	p := &Sample{Name: "Jhone Doe", Old: 20, Explain: "aaaa aaaa"}
	AppendJson(filePath, p)
}

// まるごと書き換え (既存)
func BenchmarkAppend_Overwrite(b *testing.B) {
	b.ResetTimer()
	p := &Sample{Name: "Jhone Doe", Old: 20, Explain: "aaaa aaaa"}
	OverWrite(filePath2, p)
}

// 要素を追加 (新規)
func BenchmarkAppend_AppendJsonNew(b *testing.B) {
	b.ResetTimer()
	p := &Sample{Name: "Jhone Doe", Old: 20, Explain: "aaaa aaaa"}
	AppendJson(noFilePath, p)
}

// まるごと書き換え (新規)
func BenchmarkAppend_OverwriteNew(b *testing.B) {
	b.ResetTimer()
	p := &Sample{Name: "Jhone Doe", Old: 20, Explain: "aaaa aaaa"}
	OverWrite(noFilePath2, p)
}

// delete file
func BenchmarkAppend_DeleteFile(b *testing.B) {
	os.Remove(noFilePath)
	os.Remove(noFilePath2)
}

const filePath = "./sample.json"
const noFilePath = "./no.json"
const filePath2 = "./sample2.json"
const noFilePath2 = "./no2.json"

type Sample struct {
	Name    string `json:"name"`
	Old     int    `json:"old,omitempty"`
	Explain string `json:"explain"`
}

func AppendJson(fileName string, object interface{}) {
	file, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0600)
	defer file.Close()
	fi, _ := file.Stat()
	leng := fi.Size()

	json_, _ := json.Marshal(object)

	if leng == 0 {
		file.Write([]byte(fmt.Sprintf(`[%s]`, json_)))
	} else {
		file.WriteAt([]byte(fmt.Sprintf(`,%s]`, json_)), leng-1)
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

// result
/******************************
BenchmarkAppend_AppendJson-8            1000000000               0.001882 ns/op        0 B/op          0 allocs/op
BenchmarkAppend_Overwrite-8             1000000000               0.004287 ns/op        0 B/op          0 allocs/op
BenchmarkAppend_AppendJsonNew-8         1000000000               0.002100 ns/op        0 B/op          0 allocs/op
BenchmarkAppend_OverwriteNew-8          1000000000               0.003463 ns/op        0 B/op          0 allocs/op
******************************/
