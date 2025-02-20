package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Page struct {
	after []int
	page  int
}

func main() {
	// 读取文件
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	pages := make(map[int]Page)
	var order [][]int

	// Read file line by line using bufio.Scanner
	scanner := bufio.NewScanner(file)
	emptyLineFound := false
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			emptyLineFound = true
			continue
		}
		if !emptyLineFound {
			// 空行之后的行添加到 afterEmptyLine
			re := regexp.MustCompile(`(\d+)\|(\d+)`)
			matches := re.FindStringSubmatch(line)

			if len(matches) == 3 {
				before, _ := strconv.Atoi(matches[1])
				after, _ := strconv.Atoi(matches[2])

				if page, exists := pages[before]; exists {
					// 如果存在，更新 Page 结构体中的 after 字段
					page.after = append(page.after, after)
					pages[before] = page
				} else {
					// 如果不存在，创建新的 Page 并插入到 map 中
					pages[before] = Page{after: []int{after}, page: before}
				}

			} else {
				fmt.Println("No match found.")
			}

		} else {
			// 空行之前的行添加到 beforeEmptyLine
			var row []int
			for _, str := range strings.Split(line, ",") {
				num, _ := strconv.Atoi(str) // 忽略错误处理
				row = append(row, num)
			}
			order = append(order, row)
		}
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	// 调用 part1 函数处理文件内容
	result1 := part1(pages, order)

	result2 := part2(pages, order)

	// 输出结果
	fmt.Println("Result of part1:", result1)
	fmt.Println("Result of part2:", result2)
}

func part1(pages map[int]Page, order [][]int) int {
	res := 0

	for _, ord := range order {
		// 检查当前行的顺序是否正确
		res += checkOrder(pages, ord)
	}

	return res
}

func part2(pages map[int]Page, order [][]int) int {
	res := 0

	for _, ord := range order {
		// 检查当前行的顺序是否正确
		if checkOrder(pages, ord) == 0 {
			res += resetOrder(pages, ord)
		}
	}

	return res
}

func resetOrder(pages map[int]Page, order []int) int {
	resSlice := make([]int, len(order))
	for i := 0; i < len(order); i++ {
		page := pages[order[i]]
		includes := 0
		for j := 0; j < len(order); j++ {
			if contains(page.after, order[j]) {
				includes++
			}
		}
		resSlice[len(order)-1-includes] = order[i]
	}
	return resSlice[len(resSlice)/2]
}

func checkOrder(pages map[int]Page, order []int) int {
	if order == nil || len(order) == 1 {
		return 0
	}
	for i := 0; i < len(order)-1; i++ {
		page := pages[order[i]]
		if !checkRow(&page, order[i+1:]) {
			return 0
		}
	}
	return order[len(order)/2]
}

func checkRow(page *Page, order []int) bool {
	for i := 0; i < len(order); i++ {
		if !contains(page.after, order[i]) {
			return false
		}
	}

	return true
}

func contains(arr []int, num int) bool {
	for _, val := range arr {
		if val == num {
			return true
		}
	}
	return false
}
