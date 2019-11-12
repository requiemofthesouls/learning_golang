package main

import "fmt"

// A function that takes an integer as input,
// and returns the number of bits that are equal to one in the binary representation of that number.
// You can guarantee that input is non-negative.
//
//Example: The binary representation of 1234 is 10011010010,
// so the function should return 5 in this case.
func CountBits(shit uint) int {
	sum := 0
	for i := 0; shit > 0; i++ {
		res := shit / 2
		rem := shit % 2
		if rem != 0 {
			sum += 1
		}
		shit = res
	}
	return sum
}

func main() {
	fmt.Println(CountBits(12312))
}
