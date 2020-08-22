package misc

import (
	"fmt"
)

var myMat [][]rune
var visited [][]bool

type charPos []int

var neighbours map[rune][]charPos

var rowPositionCal []int
var colPositionCal []int

func constructMatrix(listOfString []string) {
	for _, word := range listOfString {
		runeArray := []rune(word)
		myMat = append(myMat, runeArray)
		v := []bool{}
		for i := 0; i < len(word); i++ {
			v = append(v, false)
		}
		visited = append(visited, v)
	}
}

func findAdjacentOfCh(ch rune) bool {
	row, col := findPositionInMat(ch)
	if row == -1 {
		return false
	}
	visited[row][col] = true
	for index := range rowPositionCal {
		rpos := row + rowPositionCal[index]
		cpos := col + colPositionCal[index]
		if rpos == -1 || cpos == -1 || rpos == len(myMat) || cpos == len(myMat[0]) {
			// invalid coordinates
			continue
		}
		if !visited[rpos][cpos] {
			visited[rpos][cpos] = true
			pos := charPos{}
			pos = append(pos, rpos)
			pos = append(pos, cpos)
			fmt.Println("Appending ", rpos, cpos, string(myMat[rpos][cpos]))
			neighbours[ch] = append(neighbours[ch], pos)
		}
	}
	fmt.Println("neighbours of ch", string(ch), neighbours[ch])
	return true
}

func findPositionInMat(ch rune) (int, int) {
	for r := range myMat {
		for c := range myMat[r] {
			if myMat[r][c] == ch && visited[r][c] == false {
				return r, c
			}
		}
	}
	return -1, -1
}

func findWord(word []rune) bool {
	for index := 0; index < len(word)-1; index++ {
		ch := word[index]
		found := false
		if !findAdjacentOfCh(ch) {
			return false
		}
		for _, n := range neighbours[ch] {
			r := n[0]
			c := n[1]
			if word[index+1] == myMat[r][c] {
				found = true
				visited[r][c] = false
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func Run() {
	constructMatrix([]string{"DACD", "BOOK", "AIGP", "TTQR"})
	//constructMatrix([]string{"ABCD", "EFGH", "IJKL", "MNOP"})
	for _, row := range myMat {
		fmt.Println(row)
	}
	for _, row := range visited {
		fmt.Println(row)
	}
	rowPositionCal = []int{0, 1, 1, 1, 0, -1, -1, -1}
	colPositionCal = []int{1, 1, 0, -1, -1, -1, 0, 1}
	neighbours = make(map[rune][]charPos)
	fmt.Println(findWord([]rune("BAI")))
	fmt.Println("Final neigbours")
	for k := range neighbours {
		fmt.Println("neighbours of ", string(k), neighbours[k])
	}
}
