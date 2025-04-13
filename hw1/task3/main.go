package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	strSlice := []string{"ksd", "hello", "93&"}
	result1 := simpleConcatenation(strSlice)
	fmt.Println(result1)
	result2 := stringsBuilderConcatenation(strSlice)
	fmt.Println(result2)
	result3 := joinConcatenation(strSlice)
	fmt.Println(result3)
}

func stringsBuilderConcatenation(sl []string) string {
	var b strings.Builder
	for _, v := range sl {
		b.Write([]byte(v))
	}
	return b.String()
}

func simpleConcatenation(sl []string) string {
	var res string
	for _, v := range sl {
		res += v
	}
	return res
}

func joinConcatenation(sl []string) string {
	return strings.Join(sl, "")
}
