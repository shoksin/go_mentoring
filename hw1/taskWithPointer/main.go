package main

import "fmt"

func doWork(x *int) {
	y := 5
	x = &y
}

func main() {
	a := 2
	fmt.Println(a)
	doWork(&a)
	fmt.Println(a)
}

/*
    0x01	   	0x02       	 0x03	   		   0x04				         0x93	   	 0x94
|	 a		|			|	x - 0x01	| x(копия) - 0x93 |  ......  |	   y	|			|

&a = 0x01

x = ссылается на 0x01 (лежит в 0х03)

&y =  0x93

x(копия) = &y  ссылается на 0x93 (лежит в 0x93)

*/
