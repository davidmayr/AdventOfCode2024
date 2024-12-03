package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lists, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	text := string(lists)

	//Regex: (mul\((\d+),(\d+)\))
	r, _ := regexp.Compile("(mul\\(\\d+,\\d+\\))")

	response := r.FindAllString(text, -1)

	result := 0

	for _, data := range response {
		num1, num2 := parseMul(data)

		result += num1 * num2
	}

	println(result)

	//Task 2

	r2, _ := regexp.Compile("(mul\\((\\d+),(\\d+)\\))|(do\\(\\))|(don't\\(\\))")
	result2 := 0
	response2 := r2.FindAllString(text, -1)

	enabled := true

	for _, data := range response2 {
		if data == "do()" {
			enabled = true
		} else if data == "don't()" {
			enabled = false
		} else {
			if !enabled {
				continue
			}

			num1, num2 := parseMul(data)

			result2 += num1 * num2
		}
	}

	println(result2)
}

func parseMul(data string) (int, int) {
	op := strings.Replace(strings.Replace(data, "mul(", "", -1), ")", "", -1)
	numbers := strings.Split(op, ",")

	num1, err := strconv.Atoi(numbers[0])
	num2, err2 := strconv.Atoi(numbers[1])

	if err != nil || err2 != nil {
		panic(err)
	}

	return num1, num2
}
