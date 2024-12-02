package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	correct := 0
	correctTwo := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, " ")
		numbers := make([]int, 0, len(split))

		for _, s := range split {
			number, err := strconv.Atoi(s)
			if err != nil {
				fmt.Printf("Error converting string %s to int\n", s)
				panic(err)
			}
			numbers = append(numbers, number)
		}

		increasing := isIncreasing(numbers)
		decreasing := isDecreasing(numbers)
		test1 := test(numbers, increasing)

		if (increasing || decreasing) && test1 {
			correct = correct + 1
		}

		//Part two
		for i := range len(numbers) {

			listWithoutIndex := append([]int{}, numbers[:i]...)
			listWithoutIndex = append(listWithoutIndex, numbers[i+1:]...)

			increasing2 := isIncreasing(listWithoutIndex)
			decreasing2 := isDecreasing(listWithoutIndex)
			test2 := test(listWithoutIndex, increasing2)

			if (increasing2 || decreasing2) && test2 {
				correctTwo = correctTwo + 1
				break
			}
		}
	}
	println(correct)
	println(correctTwo)
}

func test(numbers []int, increasing bool) bool {
	prev := 0

	for i, number := range numbers {
		var diff int

		if increasing {
			diff = number - prev
		} else {
			diff = prev - number
		}

		if i > 0 && (diff <= 0 || diff > 3) {
			return false
		}

		prev = number
	}
	return true
}

func isIncreasing(numbers []int) bool {

	for i := range numbers {
		if i == len(numbers)-1 {
			break
		}

		if numbers[i] > numbers[i+1] {
			return false
		}
	}

	return true
}

func isDecreasing(numbers []int) bool {
	for i := range numbers {
		if i == len(numbers)-1 {
			break
		}

		if numbers[i] < numbers[i+1] {
			return false
		}
	}
	return true
}
