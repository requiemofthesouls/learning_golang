package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	dateString := "10:56"
	t, err := time.Parse("15:04", dateString)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(t)

}
