package main

import (
	"fmt"
	"strings"
)

func main() {
	findFirstStringInBracket("(hello world!)")
	anotherFindFirstStringInBracket("(hello world!)")
}

func findFirstStringInBracket(str string) string {
	indexFirstBracketFound := strings.Index(str,"(")
	indexClosingBracketFound := strings.Index(str,")")

	if indexFirstBracketFound < 0 && indexClosingBracketFound < 0{
		return ""
	}

	runes := []rune(str)
	fmt.Println(string(runes[indexFirstBracketFound+1:indexClosingBracketFound]))
	return string(runes[indexFirstBracketFound+1:indexClosingBracketFound])
}

func anotherFindFirstStringInBracket(str string) string {
	out := strings.TrimLeft(strings.TrimRight(str,")"),"(")
	fmt.Println(out)
	return out
}