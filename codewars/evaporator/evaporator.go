package main

import "fmt"

//This program tests the life of an evaporator containing a gas.

//We know the content of the evaporator (content in ml),
// the percentage of foam or gas lost every day (evap_per_day)
// and the threshold (threshold) in percentage beyond which the evaporator
// is no longer useful. All numbers are strictly positive.

//The program reports the nth day (as an integer) on which the evaporator will be out of use.

//Note : Content is in fact not necessary in the body of the function "evaporator",
// you can use it or not use it, as you wish. Some people might prefer to reason with content,
// some other with percentages only.

// It's up to you but you must keep it as a parameter because the tests have it as an argument.
func Evaporator(content float64, evapPerDay int, threshold int) int {
	// your code
	var days int
	thresholdMl := content / 100 * float64(threshold)

	for content > thresholdMl {
		amount := content / 100 * float64(evapPerDay)
		content -= amount
		days += 1
	}

	return days
}

func main() {
	fmt.Println(Evaporator(10, 10, 10))
}
