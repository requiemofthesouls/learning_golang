package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	arguments := os.Args
	var total float64 = 0

	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			total += n
		} else {
			fmt.Printf("Argument <%s> should be float.\n", arguments[i])
		}
	}

	avg := total / float64(len(arguments)-1)

	fmt.Println("Average:", avg)
}
