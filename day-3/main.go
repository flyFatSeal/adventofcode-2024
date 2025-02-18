package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// part1 处理读取的文件内容，匹配并计算结果
func part1(content string) int {
	res := 0

	// 编译正则表达式，匹配 mul(数字,数字) 形式
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// 查找所有匹配项
	matches := re.FindAllStringSubmatch(content, -1)

	// 遍历所有匹配项并处理
	for _, match := range matches {
		if len(match) == 3 {
			// 使用 strconv.Atoi 来转换字符串为整数
			val1, err1 := strconv.Atoi(match[1])
			val2, err2 := strconv.Atoi(match[2])

			if err1 == nil && err2 == nil {
				res += val1 * val2 // 按照题意进行计算，示例为乘法
			}
		}
	}

	return res
}

// part2 处理读取的文件内容，匹配并计算结果
func part2(content string) int {
	res := 0

	mulEnabled := true

	// 编译正则表达式，匹配 mul(数字,数字) 形式
	// re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)`)

	// 查找所有的匹配项
	matches := re.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		// 遇到 do() 启用 mul 指令
		if match[0] == "do()" {
			mulEnabled = true
		} else if match[0] == "don't()" {
			// 遇到 don't() 禁用 mul 指令
			mulEnabled = false
		} else if strings.HasPrefix(match[0], "mul") && mulEnabled {
			// 处理 mul(a, b) 指令，计算乘积
			val1, err1 := strconv.Atoi(match[1])
			val2, err2 := strconv.Atoi(match[2])

			// 如果启用了 mul，并且转换成功，则进行乘法计算
			if err1 == nil && err2 == nil {
				res += val1 * val2
			}
		}
	}

	return res
}

func main() {
	// 读取文件
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 将字节数据转换为字符串
	content := string(data)

	// 调用 part1 函数处理文件内容
	result1 := part1(content)

	result2 := part2(content)

	// 输出结果
	fmt.Println("Result of part1:", result1)
	fmt.Println("Result of part2:", result2)
}
