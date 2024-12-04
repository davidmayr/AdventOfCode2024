package main

import (
	"os"
	"strings"
)

func main() {
	lists, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	text := string(lists)
	lines := strings.Split(text, "\n")

	//Part 1

	//Y to X 2d array
	var data = make([][]string, len(lines))

	for i, line := range lines {
		data[i] = strings.Split(line, "")
	}

	counter := 0

	for il, line := range data {

		for ic, _ := range line {

			if isThere(data, ic, il, 0, 1) {
				counter++
			}
			if isThere(data, ic, il, 1, 1) {
				counter++
			}
			if isThere(data, ic, il, 1, 0) {
				counter++
			}
			if isThere(data, ic, il, 1, -1) {
				counter++
			}
			if isThere(data, ic, il, 0, -1) {
				counter++
			}
			if isThere(data, ic, il, -1, -1) {
				counter++
			}
			if isThere(data, ic, il, -1, 0) {
				counter++
			}
			if isThere(data, ic, il, -1, 1) {
				counter++
			}

		}
	}

	println(counter)

	//Part two

	counterTwo := 0

	for y, line := range data {

		for x, _ := range line {

			char, _ := getChar(data, x, y)

			if char != "A" {
				continue
			}

			/*
			 * S S
			 *  A
			 * M M
			 */
			if isThere2(data, x-1, y+1, 1, -1) && isThere2(data, x+1, y+1, -1, -1) {
				counterTwo++
			}

			/*
			 * M S
			 *  A
			 * M S
			 */
			if isThere2(data, x-1, y-1, 1, 1) && isThere2(data, x-1, y+1, 1, -1) {
				counterTwo++
			}

			/*
			 * M M
			 *  A
			 * S S
			 */
			if isThere2(data, x-1, y-1, 1, 1) && isThere2(data, x+1, y-1, -1, 1) {
				counterTwo++
			}

			/*
			 * S M
			 *  A
			 * S M
			 */
			if isThere2(data, x+1, y-1, -1, 1) && isThere2(data, x+1, y+1, -1, -1) {
				counterTwo++
			}

		}
	}

	println(counterTwo)
}

func isThere(data [][]string, startPosX int, startPosY int, modX int, modY int) bool {
	return isThereChar(data, startPosX, startPosY, modX, modY, "X", 0) &&
		isThereChar(data, startPosX, startPosY, modX, modY, "M", 1) &&
		isThereChar(data, startPosX, startPosY, modX, modY, "A", 2) &&
		isThereChar(data, startPosX, startPosY, modX, modY, "S", 3)
}

func isThere2(data [][]string, startPosX int, startPosY int, modX int, modY int) bool {
	return isThereChar(data, startPosX, startPosY, modX, modY, "M", 0) &&
		isThereChar(data, startPosX, startPosY, modX, modY, "A", 1) &&
		isThereChar(data, startPosX, startPosY, modX, modY, "S", 2)
}

func isThereChar(data [][]string, startPosX int, startPosY int, modX int, modY int, char string, index int) bool {
	y := startPosY + (modY * index)
	x := startPosX + (modX * index)

	t, succ := getChar(data, x, y)

	if !succ {
		return false
	}

	return t == char
}

func getChar(data [][]string, x int, y int) (string, bool) {

	if y < 0 || y >= len(data) {
		return "", false
	}

	yD := data[y]

	if x < 0 || x >= len(yD) {
		return "", false
	}

	return yD[x], true
}
