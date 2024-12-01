package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lists, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	text := string(lists)
	lines := strings.Split(text, "\n")

	l1 := make([]int, 0, len(lines))
	l2 := make([]int, 0, len(lines))

	for i := range lines {
		if len(lines[i]) == 0 {
			continue
		}

		singleLine := strings.Split(lines[i], "   ")

		l1Parse, errorl1 := strconv.Atoi(singleLine[0])
		if errorl1 != nil {
			panic(errorl1)
		}

		l2Parse, errorl2 := strconv.Atoi(singleLine[1])
		if errorl2 != nil {
			panic(errorl2)
		}

		l1 = append(l1, l1Parse)
		l2 = append(l2, l2Parse)
	}

	sort.Ints(l1)
	sort.Ints(l2)

	result := 0

	for i := range l1 {
		number := l1[i]
		otherNumber := l2[i]

		if number < otherNumber {
			number, otherNumber = otherNumber, number
		}

		result += number - otherNumber
	}
	println("The result is: " + strconv.FormatInt(int64(result), 10))

	//Part two

	findings := map[int]int{}

	for i := range l2 {
		value := l2[i]

		val, _ := findings[value]
		findings[value] = val + 1
	}

	resultTwo := 0

	for i := range l1 {
		value := l1[i]

		findings, _ := findings[value]

		resultTwo += value * findings
	}

	println(resultTwo)
}
