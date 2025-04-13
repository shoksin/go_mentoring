package main

import (
	"testing"
)

func BenchmarkBuilderConcatenation(b *testing.B) {
	strSlice := make([]string, 1000)
	for i := 0; i < len(strSlice); i++ {
		strSlice[i] = "example"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stringsBuilderConcatenation(strSlice)
	}
}

func BenchmarkSimpleConcatenation(b *testing.B) {
	strSlice := make([]string, 1000)
	for i := 0; i < len(strSlice); i++ {
		strSlice[i] = "example"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		simpleConcatenation(strSlice)
	}
}

func BenchmarkJoinConcatenation(b *testing.B) {
	strSlice := make([]string, 1000)
	for i := 0; i < len(strSlice); i++ {
		strSlice[i] = "example"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		joinConcatenation(strSlice)
	}
}
