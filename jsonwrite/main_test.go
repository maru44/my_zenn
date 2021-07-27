package main

import (
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
