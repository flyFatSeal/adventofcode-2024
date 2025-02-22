package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Page struct {
	after []int
	page  int
}

type Calc struct {
	sum   int
	items []int
}

func main() {
	// 读取文件

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read file line by line using bufio.Scanner
	scanner := bufio.NewScanner(file)

	calcs := make([]Calc, 0)

	for scanner.Scan() {
		line := scanner.Text()

		// 空行之后的行添加到 afterEmptyLine
		re := regexp.MustCompile(`\d+`)

		// 查找所有匹配的数字
		matches := re.FindAllString(line, -1)

		calc := Calc{}
		for i := 0; i < len(matches); i++ {
			if i == 0 {
				calc.sum, _ = strconv.Atoi(matches[i])
			} else {
				val, _ := strconv.Atoi(matches[i])
				calc.items = append(calc.items, val)
			}

		}
		calcs = append(calcs, calc)

	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	// 创建一个新的切片来复制原始切片
	copySlice := make([]Calc, len(calcs))

	// 使用 copy 函数进行复制
	copy(copySlice, calcs)

	// 调用 part1 函数处理文件内容
	result1 := part1(calcs)
	result2 := part2(copySlice)
	// 输出结果
	fmt.Println("Result of part1:", result1)
	fmt.Println("Result of part2:", result2)
}

func part1(calcs []Calc) int {
	res := 0
	for _, calc := range calcs {
		sums := calcSum([]int{}, reverse(calc.items))
		for _, sum := range sums {
			if sum == calc.sum {
				res += calc.sum
				break
			}
		}
	}

	return res
}

func part2(calcs []Calc) int {
	res := 0
	for _, calc := range calcs {
		sums := calcSum2([]int{}, reverse(calc.items))
		for _, sum := range sums {
			if sum == calc.sum {
				res += calc.sum
				break
			}
		}
	}

	return res
}

func calcSum(res []int, items []int) []int {
	if len(items) == 0 {
		return []int{0}
	}

	if len(items) == 1 {
		return []int{items[0]}
	}

	itemSums := calcSum(res, items[1:])

	for _, sum := range itemSums {
		res = append(res, sum+items[0])
		res = append(res, sum*items[0])
	}

	return res
}

func calcSum2(res []int, items []int) []int {
	if len(items) == 0 {
		return []int{0}
	}

	if len(items) == 1 {
		return []int{items[0]}
	}

	itemSums := calcSum(res, items[1:])

	for _, sum := range itemSums {
		res = append(res, sum+items[0])
		res = append(res, sum*items[0])
		res = append(res, concatenate(sum, items[0]))
	}

	return res
}

func reverse(arr []int) []int {
	// 创建一个新的切片存储反转后的元素
	result := make([]int, len(arr))
	for i, v := range arr {
		result[len(arr)-1-i] = v
	}
	return result
}

func concatenate(a, b int) int {
	// 通过左移和按位或操作拼接数字
	shift := 0
	temp := b
	for temp > 0 {
		shift++
		temp /= 10
	}
	return a*intPow(10, shift) + b
}

// 计算 10 的幂
func intPow(x, n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= x
	}
	return result
}
