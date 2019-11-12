package main

import "fmt"

func PlayPass(s string, n int) string {
	for index, letter := range s {
		fmt.Printf("%d %c", index, letter)
	}
	return "s"
}

func main() {
	fmt.Println(PlayPass("loool", 2))
}
