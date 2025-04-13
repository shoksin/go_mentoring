package main

import "fmt"

func main() {
	testStrings := []string{"Hello 2)", "ksdщеD", "漢6字hi紅"}

	for _, v := range testStrings {

		fmt.Println(v, "using iterating through runes:")

		for i, ch := range v {
			fmt.Printf("pos %d - char %c\n", i+1, ch)
		}

		fmt.Println(v, "using iterating through bytes:")

		for i := 0; i < len(v); i++ {
			fmt.Printf("pos %d - char %c\n", i+1, v[i])
		}
		fmt.Println()
	}

}
