package arrays

import (
	"fmt"
	"strings"
)

func printAllSubsets(input, output []rune) {
	if len(input) == 0 {
		fmt.Printf(string(output) + ", ")
		return
	}

	// case the character is not taken
	firstChar := input[0]
	newInput := input[1:]
	// output remains the same
	printAllSubsets(newInput, output)
	// case when the character is choosen for the set
	output = append(output, firstChar)
	printAllSubsets(newInput, output)

}

func addspaceBefore(ch rune, input []rune) []rune {
	output := []rune{}
	for _, chs := range input {
		if chs == ch {
			output = append(output, '_')
		}
		output = append(output, chs)

	}
	return output
}

func permutationsOfSpaces(input, output []rune) {
	if len(input) == 0 {
		fmt.Println(string(output))
		return
	}

	// case the character is not taken
	firstChar := input[0]
	newInput := input[1:]
	// output remains the same
	output = append(output, firstChar)
	permutationsOfSpaces(newInput, output)
	output = addspaceBefore(firstChar, output)
	//output = append(output, firstChar)
	// case when the character is choosen for the set
	permutationsOfSpaces(newInput, output)
}

var mp map[string]bool

var output1, output2 []rune

func convertToUpper(ch rune) rune {
	// var oldoutput []rune
	// for index, chars := range output {
	// 	oldoutput = append(oldoutput, chars)
	// 	if chars == ch {
	// 		uc := strings.ToUpper(string(ch))
	// 		output[index] = []rune(uc)[0]
	// 		//break
	// 	}
	// }
	// fmt.Printf("converting %v to %v\n", string(oldoutput), string(output))
	// return output
	uc := strings.ToUpper(string(ch))
	return []rune(uc)[0]
}

func permutationsOfCharCases(input, output []rune) {
	if len(input) == 0 {
		//fmt.Println(string(output))
		if _, ok := mp[string(output)]; !ok {
			fmt.Println(string(output))
			mp[string(output)] = true
		}
		return
	}
	// output remains the same
	firstchar := input[0]
	output1 = append(output, firstchar)
	permutationsOfCharCases(input[1:], output1)
	output2 = append(output, convertToUpper(firstchar))
	permutationsOfCharCases(input[1:], output2)
}

// func Run() {
// 	input := []rune("a1b2c")
// 	mp = make(map[string]bool)
// 	//permutationsOfSpaces(input[1:], []rune("a"))
// 	permutationsOfCharCases(input, []rune(""))
// 	fmt.Println(mp)
// }
