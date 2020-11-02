package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	dateString := "28 November 1999"
	t, err := time.Parse("02 January 2006", dateString)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(t)

}
