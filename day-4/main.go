package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isMatchMAS(ceil []string) bool {
	for index, cell := range ceil {
		if index == 0 && cell != "M" {
			return false
		}

		if index == 1 && cell != "A" {
			return false
		}

		if index == 2 && cell != "S" {
			return false
		}

	}
	return true
}

func countXMASPath(location []int, inputs [][]string) int {
	res := 0
	maxX := len(inputs[0])
	maxY := len(inputs)

	x, y := location[0], location[1]
	// 正横向匹配
	if x <= maxX-4 {
		ceils := []string{inputs[y][x+1], inputs[y][x+2], inputs[y][x+3]}
		if isMatchMAS(ceils) {
			res++
		}
	}

	// 负横向匹配
	if x >= 3 {
		ceils := []string{inputs[y][x-1], inputs[y][x-2], inputs[y][x-3]}
		if isMatchMAS(ceils) {
			res++
		}
	}

	// 向下匹配
	if y <= maxY-4 {
		ceils := []string{inputs[y+1][x], inputs[y+2][x], inputs[y+3][x]}
		if isMatchMAS(ceils) {
			res++
		}
	}

	// 向上匹配
	if y >= 3 {
		ceils := []string{inputs[y-1][x], inputs[y-2][x], inputs[y-3][x]}
		if isMatchMAS(ceils) {
			res++
		}
	}

	// 正斜向匹配
	if x <= maxX-4 && y <= maxY-4 {
		ceils := []string{inputs[y+1][x+1], inputs[y+2][x+2], inputs[y+3][x+3]}
		if isMatchMAS(ceils) {
			res++
		}
	}
	// 负斜向匹配
	if x >= 3 && y >= 3 {
		ceils := []string{inputs[y-1][x-1], inputs[y-2][x-2], inputs[y-3][x-3]}
		if isMatchMAS(ceils) {
			res++
		}
	}
	// 反斜向匹配
	if x <= maxX-4 && y >= 3 {
		ceils := []string{inputs[y-1][x+1], inputs[y-2][x+2], inputs[y-3][x+3]}
		if isMatchMAS(ceils) {
			res++
		}
	}
	// 反斜向匹配
	if x >= 3 && y <= maxY-4 {
		ceils := []string{inputs[y+1][x-1], inputs[y+2][x-2], inputs[y+3][x-3]}
		if isMatchMAS(ceils) {
			res++
		}
	}

	return res
}

func isMAXPath(location []int, inputs [][]string) bool {
	res := 0

	maxX := len(inputs[0])
	maxY := len(inputs)

	x, y := location[0], location[1]

	if x == 0 || x == maxX-1 || y == 0 || y == maxY-1 {
		return false
	}

	if inputs[y-1][x-1] == "M" {
		if inputs[y+1][x+1] == "S" {
			res++
		}
	} else if inputs[y-1][x-1] == "S" {
		if inputs[y+1][x+1] == "M" {
			res++
		}
	}

	if inputs[y-1][x+1] == "M" {
		if inputs[y+1][x-1] == "S" {
			res++
		}
	} else if inputs[y-1][x+1] == "S" {
		if inputs[y+1][x-1] == "M" {
			res++
		}
	}

	return res == 2
}

// part1 处理读取的文件内容，匹配并计算结果
func part1(inputs [][]string) int {
	res := 0

	for y, row := range inputs {
		for x, cell := range row {
			if cell == "X" {
				res += countXMASPath([]int{x, y}, inputs)
			}
		}
	}

	return res
}

func part2(inputs [][]string) int {
	res := 0

	for y, row := range inputs {
		for x, cell := range row {
			if cell == "A" && isMAXPath([]int{x, y}, inputs) {
				res += 1
			}
		}
	}

	return res
}

func main() {
	// 读取文件
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var inputs [][]string

	// Read file line by line using bufio.Scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		inputs = append(inputs, strings.Split(line, ""))
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	// 调用 part1 函数处理文件内容
	result1 := part1(inputs)

	result2 := part2(inputs)

	// 输出结果
	fmt.Println("Result of part1:", result1)
	fmt.Println("Result of part2:", result2)
}
