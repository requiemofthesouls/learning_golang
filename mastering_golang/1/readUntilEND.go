package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f := os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	fmt.Println("Please give me a real numbers. Type 'END' for cancel.")

	for scanner.Scan() {
		if scanner.Text() != "END" {
			n, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				fmt.Printf("<%s> is not a valid integer.\n", scanner.Text())
			} else {
				fmt.Println(">", n)
			}
		} else {
			fmt.Println("Got 'END', returning...")
			return
		}
	}

}
