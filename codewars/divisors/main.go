package main

import "fmt"

func Divisors(n int) int {
	divisors := 0

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divisors++
		}
	}

	return divisors
}

func main() {
	fmt.Println(Divisors(10))
}
