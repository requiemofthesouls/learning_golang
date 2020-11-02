package main

import (
	"fmt"
)

func main() {

	a := [10]string{"test1", "test2", "test3", "test4", "test5", "test6", "test7", "test8", "test9", "test10"}
	fmt.Println(a)

	m := make(map[int]string)
	for index, value := range a {
		m[index+1] = value
	}

	fmt.Println(m)
}
