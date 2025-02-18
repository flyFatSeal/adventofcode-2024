package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// Code here
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	var column1 []int
	var column2 []int

	// 使用 bufio.Scanner 逐行读取文件
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 获取当前行
		line := scanner.Text()

		// 将当前行按空格或制表符分割成多个字段
		parts := strings.Fields(line)

		// 确保每行有两列数据
		if len(parts) == 2 {
			// 将每列的数据转换为整数并添加到相应的切片中
			var val1, val2 int
			_, err := fmt.Sscanf(parts[0], "%d", &val1)
			if err != nil {
				fmt.Println("Error parsing first column:", err)
				continue
			}
			_, err = fmt.Sscanf(parts[1], "%d", &val2)
			if err != nil {
				fmt.Println("Error parsing second column:", err)
				continue
			}

			// 将转换后的值添加到切片中
			column1 = append(column1, val1)
			column2 = append(column2, val2)
		}
	}
	// 检查是否有扫描错误
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	sort.Ints(column1)
	sort.Ints(column2)

	res := 0

	for i := 0; i < len(column1); i++ {
		res += absInt(column1[i] - column2[i])
	}

	fmt.Println("The result is:", res)
}

func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
