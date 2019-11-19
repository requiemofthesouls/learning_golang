package main

import "math"

// (a ^ p + b ^ (p+1) + c ^(p+2) + d ^ (p+3) + ...) = n * k
func DigPow(n, p int) int {
	nLen := int(math.Round(math.Log10(float64(n))))
	digitArray := make([]int, nLen)
	d := n
	for i := nLen - 1; i >= 0; i -= 1 {
		digitArray[i] = d % 10
		d /= 10
	}

	var sum int
	for _, digit := range digitArray {
		sum += int(math.Pow(float64(digit), float64(p)))
		p += 1
	}

	if sum/n >= 1 {
		return sum / n
	}

	return -1
}
