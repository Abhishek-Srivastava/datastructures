package str

import (
	"fmt"
	"strings"
)

func reverseWord(s string) string {
	reverseWord := strings.Builder{}
	for i, _ := range s {
		reverseWord.WriteByte(s[len(s)-i-1])
	}
	return reverseWord.String()
}

func reverse(s string) string {
	words := strings.Split(s, " ")
	reversedString := strings.Builder{}
	for i, _ := range words {
		reversedString.WriteString(words[len(words)-1-i])
		reversedString.WriteString(" ")
	}
	return reversedString.String()
}
func Run() {
	fmt.Println(reverse("I am vengeance"))
}
