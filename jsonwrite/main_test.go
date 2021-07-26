package main

import (
	"testing"
)

// 要素を追加
func BenchmarkAppend_AppendJson(b *testing.B) {
	b.ResetTimer()
	AppendJson()
}

// まるごと書き換え
func BenchmarkAppend_Overwrite(b *testing.B) {
	b.ResetTimer()
	OverWrite()
}
