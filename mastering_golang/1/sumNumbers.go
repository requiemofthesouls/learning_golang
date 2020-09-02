package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	arguments := os.Args
	var total int64 = 0

	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseInt(arguments[i], 10, 64)
		if err == nil {
			total += n
		} else {
			fmt.Printf("Argument <%s> should be a real number.\n", arguments[i])
		}
	}

	fmt.Println("Total:", total)
}
