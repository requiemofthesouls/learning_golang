package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

var lowercaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
var uppercaseAlphabet = strings.ToUpper(lowercaseAlphabet)

func loopCharShifter(char rune, plan string, n int) (shiftedChar rune) {
	fmt.Printf("SHIFTING {%c}\n", char)
	shiftedChar = char
	return
}

// Takes care of symbol shifting on giving n steps
func shiftChar(char rune, n int) string {
	if unicode.IsLetter(char) {
		var alphabet string
		if unicode.IsLower(char) {
			alphabet = lowercaseAlphabet
		} else {
			alphabet = uppercaseAlphabet
		}
		currentCharIndex := strings.IndexRune(alphabet, char)
		for ; n > 0; n-- {
			if currentCharIndex <= len(alphabet) {
				char += 1
			} else {
				char = rune(alphabet[0])
			}
		}
		shiftedChar := char
		return string(shiftedChar)
	} else if unicode.IsDigit(char) {
		shiftedInt := int('9' - char)
		return strconv.Itoa(shiftedInt)
	} else {
		return string(char)
	}

}

//1. shift each letter by a given number but the transformed letter must be a letter (circular shift),
//2. replace each digit by its complement to 9,
//3. keep such as non alphabetic and non digit characters,
//4. downcase each letter in odd position, upcase each letter in even position (the first character is in position 0),
//5. reverse the whole result.
func PlayPass(phrase string, n int) (passphrase string) {
	if n == 0 {
		return passphrase
	}

	var passphraseArray []string

	for _, char := range phrase {
		shiftedChar := shiftChar(char, n)
		passphraseArray = append(passphraseArray, shiftedChar)
	}

	for i, str := range passphraseArray {
		if (i+1)%2 == 0 {
			passphraseArray[i] = strings.ToLower(str)
		} else {
			passphraseArray[i] = strings.ToUpper(str)
		}
	}

	reversedPassphraseArray := passphraseArray
	fmt.Println(strings.Join(passphraseArray, ""))
	// TODO: Почему затирает пробелы при reverse
	sort.Sort(sort.Reverse(sort.StringSlice(reversedPassphraseArray)))
	passphrase = strings.Join(passphraseArray, "")
	return
}

func main() {
	fmt.Println(PlayPass("I LOVE YOU!!!", 1))
}
