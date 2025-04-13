package main

import (
	"fmt"
	"sync"
)

func main() {
	s := "x"
	if isPalindromeSimpleCheck(s) {
		fmt.Println(s, "is polindrome.")
	} else {
		fmt.Println(s, "is not polindrome.")
	}

	if isPalindromeAsyncCheck(s) {
		fmt.Println(s, "is polindrome.")
	} else {
		fmt.Println(s, "is not polindrome.")
	}
}

func isPalindromeSimpleCheck(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isPalindromeAsyncCheck(s string) bool {
	if len(s) < 2 {
		return true
	}
	n := getN(s)
	intervals := getIntervals(len(s), n)
	var wg sync.WaitGroup
	var mu sync.Mutex
	isPolindrome := true
	for i := 0; i <= n/2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s1 := intervals[i]
			for j := s1[0]; j < s1[1]; j++ {
				if s[j] != s[len(s)-1-j] {
					mu.Lock()
					isPolindrome = false
					mu.Unlock()
					return
				}
			}
		}()
	}

	wg.Wait()
	return isPolindrome
}

func getIntervals(lenght int, n int) [][]int {
	res := make([][]int, 0)
	for i := range res {
		res[i] = make([]int, 0)
	}

	l := lenght / n

	centralElement := -1
	if lenght%2 == 1 {
		centralElement = lenght / 2
	}

	for i := 0; i < lenght; i += l {
		if i == centralElement {
			i++
			res = append(res, []int{i, i + l})
		} else {
			res = append(res, []int{i, i + l})
		}
	}

	return res
}

func getN(s string) int {
	if len(s) == 2 {
		return 2
	}
	if (len(s))%2 == 1 {
		return len(s) / 2
	}
	return (len(s) - 1) / 2
}
